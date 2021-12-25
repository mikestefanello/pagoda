package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"goweb/config"
	"goweb/funcmap"
)

// TemplateRenderer provides a flexible and easy to use method of rendering simple templates or complex sets of
// templates while also providing caching and/or hot-reloading depending on your current environment
type TemplateRenderer struct {
	// templateCache stores a cache of parsed page templates
	templateCache sync.Map

	// funcMap stores the template function map
	funcMap template.FuncMap

	// templatePath stores the complete path to the templates directory
	templatesPath string

	// config stores application configuration
	config *config.Config
}

// NewTemplateRenderer creates a new TemplateRenderer
func NewTemplateRenderer(cfg *config.Config) *TemplateRenderer {
	t := &TemplateRenderer{
		templateCache: sync.Map{},
		funcMap:       funcmap.GetFuncMap(),
		config:        cfg,
	}

	// Gets the complete templates directory path
	// This is needed in case this is called from a package outside of main, such as within tests
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	t.templatesPath = filepath.Join(filepath.Dir(d), config.TemplateDir)

	return t
}

// Parse parses a set of templates and caches them for quick execution
// If the application environment is set to local, the cache will be bypassed and templates will be
// parsed upon each request so hot-reloading is possible without restarts.
//
// All template files and template directories must be provided relative to the templates directory
// and without template extensions. Those two values can be altered via the config package.
//
// cacheGroup is used to separate templates in to groups within the cache to avoid potential conflicts
// with the cacheID.
//
// baseName is the filename of the base template without any paths or an extension.
// files is a slice of all individual template files that will be included in the parse.
// directories is a slice of directories which all template files witin them will be included in the parse
//
// Also included will be the function map provided by the funcmap package.
//
// An example usage of this:
// t.Parse(
//		"page",
//		"home",
//		"main",
//		[]string{
//			"layouts/main",
//			"pages/home",
//		},
//		[]string{"components"},
//)
//
// This will perform a template parse which will:
//		- Be cached using a key of "page:home"
//		- Include the layouts/main.gohtml and pages/home.gohtml templates
//		- Include all templates within the components directory
//		- Include the function map within the funcmap package
//		- Set the base template as main.gohtml
func (t *TemplateRenderer) Parse(cacheGroup, cacheID, baseName string, files []string, directories []string) error {
	cacheKey := t.getCacheKey(cacheGroup, cacheID)

	// Check if the template has not yet been parsed or if the app environment is local, so that
	// templates reflect changes without having the restart the server
	if _, err := t.Load(cacheGroup, cacheID); err != nil || t.config.App.Environment == config.EnvLocal {
		// Initialize the parsed template with the function map
		parsed := template.New(baseName + config.TemplateExt).
			Funcs(t.funcMap)

		// Parse all files provided
		if len(files) > 0 {
			for k, v := range files {
				files[k] = fmt.Sprintf("%s/%s%s", t.templatesPath, v, config.TemplateExt)
			}

			parsed, err = parsed.ParseFiles(files...)
			if err != nil {
				return err
			}
		}

		// Parse all templates within the provided directories
		for _, dir := range directories {
			dir = fmt.Sprintf("%s/%s/*%s", t.templatesPath, dir, config.TemplateExt)
			parsed, err = parsed.ParseGlob(dir)
			if err != nil {
				return err
			}
		}

		// Store the template so this process only happens once
		t.templateCache.Store(cacheKey, parsed)
	}

	return nil
}

// Execute executes a cached template with the data provided
// See Parse() for an explanation of the parameters
func (t *TemplateRenderer) Execute(cacheGroup, cacheID, baseName string, data interface{}) (*bytes.Buffer, error) {
	tmpl, err := t.Load(cacheGroup, cacheID)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, baseName+config.TemplateExt, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// ParseAndExecute is a wrapper around Parse() and Execute()
func (t *TemplateRenderer) ParseAndExecute(cacheGroup, cacheID, baseName string, files []string, directories []string, data interface{}) (*bytes.Buffer, error) {
	var buf *bytes.Buffer
	var err error

	if err = t.Parse(cacheGroup, cacheID, baseName, files, directories); err != nil {
		return nil, err
	}
	if buf, err = t.Execute(cacheGroup, cacheID, baseName, data); err != nil {
		return nil, err
	}

	return buf, nil
}

// Load loads a template from the cache
func (t *TemplateRenderer) Load(cacheGroup, cacheID string) (*template.Template, error) {
	load, ok := t.templateCache.Load(t.getCacheKey(cacheGroup, cacheID))
	if !ok {
		return nil, errors.New("uncached page template requested")
	}

	tmpl, ok := load.(*template.Template)
	if !ok {
		return nil, errors.New("unable to cast cached template")
	}

	return tmpl, nil
}

// GetTemplatesPath gets the complete path to the templates directory
func (t *TemplateRenderer) GetTemplatesPath() string {
	return t.templatesPath
}

// getCacheKey gets a cache key for a given group and ID
func (t *TemplateRenderer) getCacheKey(cacheGroup, cacheID string) string {
	return fmt.Sprintf("%s:%s", cacheGroup, cacheID)
}
