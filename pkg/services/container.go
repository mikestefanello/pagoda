package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikestefanello/backlite"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/log"
	inertia "github.com/romsar/gonertia/v2"
	"github.com/spf13/afero"

	// Required by ent.
	_ "github.com/mikestefanello/pagoda/ent/runtime"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests.
type Container struct {
	// Validator stores a validator
	Validator *Validator

	// Web stores the web framework.
	Web *echo.Echo

	// Config stores the application configuration.
	Config *config.Config

	// Cache contains the cache client.
	Cache *CacheClient

	// Database stores the connection to the database.
	Database *sql.DB

	// Files stores the file system.
	Files afero.Fs

	// ORM stores a client to the ORM.
	ORM *ent.Client

	// Graph is the entity graph defined by your Ent schema.
	Graph *gen.Graph

	// Mail stores an email sending client.
	Mail *MailClient

	// Auth stores an authentication client.
	Auth *AuthClient

	// Tasks stores the task client.
	Tasks *backlite.Client

	// Inertia for React
	Inertia *inertia.Inertia
}

// NewContainer creates and initializes a new Container.
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initFiles()
	c.initORM()
	c.initAuth()
	c.initMail()
	c.initTasks()
	c.initInertia()
	return c
}

// Shutdown gracefully shuts the Container down and disconnects all connections.
func (c *Container) Shutdown() error {
	// Shutdown the web server.
	webCtx, webCancel := context.WithTimeout(context.Background(), c.Config.HTTP.ShutdownTimeout)
	defer webCancel()
	if err := c.Web.Shutdown(webCtx); err != nil {
		return err
	}

	// Shutdown the task runner.
	taskCtx, taskCancel := context.WithTimeout(context.Background(), c.Config.Tasks.ShutdownTimeout)
	defer taskCancel()
	c.Tasks.Stop(taskCtx)

	// Shutdown the ORM.
	if err := c.ORM.Close(); err != nil {
		return err
	}

	// Shutdown the database.
	if err := c.Database.Close(); err != nil {
		return err
	}

	// Shutdown the cache.
	c.Cache.Close()

	return nil
}

// initConfig initializes configuration.
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg

	// Configure logging.
	switch cfg.App.Environment {
	case config.EnvProduction:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

// initValidator initializes the validator.
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web framework.
func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true
	c.Web.Validator = c.Validator
}

// initCache initializes the cache.
func (c *Container) initCache() {
	store, err := newInMemoryCache(c.Config.Cache.Capacity)
	if err != nil {
		panic(err)
	}

	c.Cache = NewCacheClient(store)
}

// initDatabase initializes the database.
func (c *Container) initDatabase() {
	var err error
	var connection string

	switch c.Config.App.Environment {
	case config.EnvTest:
		// TODO: Drop/recreate the DB, if this isn't in memory?
		connection = c.Config.Database.TestConnection
	default:
		connection = c.Config.Database.Connection
	}

	c.Database, err = openDB(c.Config.Database.Driver, connection)
	if err != nil {
		panic(err)
	}
}

// initFiles initializes the file system.
func (c *Container) initFiles() {
	// Use in-memory storage for tests.
	if c.Config.App.Environment == config.EnvTest {
		c.Files = afero.NewMemMapFs()
		return
	}

	fs := afero.NewOsFs()
	if err := fs.MkdirAll(c.Config.Files.Directory, 0755); err != nil {
		panic(err)
	}
	c.Files = afero.NewBasePathFs(fs, c.Config.Files.Directory)
}

// initORM initializes the ORM.
func (c *Container) initORM() {
	drv := entsql.OpenDB(c.Config.Database.Driver, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	// Load the graph.
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	p := filepath.Join(filepath.Dir(d), "../ent/schema")
	g, err := entc.LoadGraph(p, &gen.Config{})
	if err != nil {
		panic(err)
	}
	c.Graph = g
}

// initAuth initializes the authentication client.
func (c *Container) initAuth() {
	c.Auth = NewAuthClient(c.Config, c.ORM)
}

// initMail initialize the mail client.
func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config)
	if err != nil {
		panic(fmt.Sprintf("failed to create mail client: %v", err))
	}
}

// initTasks initializes the task client.
func (c *Container) initTasks() {
	var err error
	// You could use a separate database for tasks, if you'd like. but using one
	// makes transaction support easier.
	c.Tasks, err = backlite.NewClient(backlite.ClientConfig{
		DB:              c.Database,
		Logger:          log.Default(),
		NumWorkers:      c.Config.Tasks.Goroutines,
		ReleaseAfter:    c.Config.Tasks.ReleaseAfter,
		CleanupInterval: c.Config.Tasks.CleanupInterval,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create task client: %v", err))
	}

	if err = c.Tasks.Install(); err != nil {
		panic(fmt.Sprintf("failed to install task schema: %v", err))
	}
}

func ProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}

func (c *Container) getInertia() *inertia.Inertia {
	rootDir := ProjectRoot()
	viteHotFile := filepath.Join(rootDir, "public", "hot")
	rootViewFile := filepath.Join(rootDir, "resources", "views", "root.html")
	manifestPath := filepath.Join(rootDir, "public", "build", "manifest.json")
	viteManifestPath := filepath.Join(rootDir, "public", "build", ".vite", "manifest.json")
	// flashProvider := provider.NewSessionFlashProvider()

	// check if laravel-vite-plugin is running in dev mode (it puts a "hot" file in the public folder)
	url, err := viteHotFileUrl(viteHotFile)
	if err != nil {
		panic(err)
	}
	if url != "" {
		i, err := inertia.NewFromFile(
			rootViewFile,
			// inertia.WithFlashProvider(flashProvider),
		)
		if err != nil {
			panic(err)
		}

		i.ShareTemplateFunc("vite", func(entry string) (template.HTML, error) {
			if entry != "" && !strings.HasPrefix(entry, "/") {
				entry = "/" + entry
			}
			htmlTag := fmt.Sprintf(`<script type="module" src="%s%s"></script>`, url, entry)
			return template.HTML(htmlTag), nil
		})
		return i
	}

	// laravel-vite-plugin not running in dev mode, use build manifest file
	// check if the manifest file exists, if not, rename it
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		// move the manifest from ./public/build/.vite/manifest.json to ./public/build/manifest.json
		// so that the vite function can find it
		err := os.Rename(viteManifestPath, manifestPath)
		if err != nil {
			return nil
		}
	}

	i, err := inertia.NewFromFile(
		rootViewFile,
		inertia.WithVersionFromFile(manifestPath),
		// inertia.WithFlashProvider(flashProvider),
	)
	if err != nil {
		panic(err)
	}

	i.ShareTemplateFunc("vite", vite(manifestPath, "/public/build/"))

	return i
}

func (c *Container) initInertia() {
	c.Inertia = c.getInertia()
}

func vite(manifestPath, buildDir string) func(path string) (string, error) {
	f, err := os.Open(manifestPath)
	if err != nil {
		log.Default().Error("cannot open provided vite manifest file: %s", err)
		panic(err)
	}
	defer f.Close()

	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})
	err = json.NewDecoder(f).Decode(&viteAssets)
	// print content of viteAssets
	for k, v := range viteAssets {
		log.Default().Debug("%s: %s\n", k, v.File)
	}

	if err != nil {
		log.Default().Error("cannot unmarshal vite manifest file to json: %s", err)
		panic(err)
	}

	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

// openDB opens a database connection.
func openDB(driver, connection string) (*sql.DB, error) {
	if driver == "sqlite3" {
		// Helper to automatically create the directories that the specified sqlite file
		// should reside in, if one.
		d := strings.Split(connection, "/")
		if len(d) > 1 {
			dirpath := strings.Join(d[:len(d)-1], "/")

			if err := os.MkdirAll(dirpath, 0755); err != nil {
				return nil, err
			}
		}

		// Check if a random value is required, which is often used for in-memory test databases.
		if strings.Contains(connection, "$RAND") {
			connection = strings.Replace(connection, "$RAND", fmt.Sprint(rand.Int()), 1)
		}
	}

	return sql.Open(driver, connection)
}

// viteHotFileUrl Get the vite hot file url
func viteHotFileUrl(viteHotFile string) (string, error) {
	_, err := os.Stat(viteHotFile)
	if err != nil {
		return "", nil
	}
	content, err := os.ReadFile(viteHotFile)
	if err != nil {
		return "", err
	}
	url := strings.TrimSpace(string(content))
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		url = url[strings.Index(url, ":")+1:]
	} else {
		url = "//localhost:1323"
	}
	return url, nil
}
