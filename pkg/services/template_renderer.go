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

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/funcmap"
)

type (
	// TemplateRenderer provides a flexible and easy to use method of rendering simple templates or complex sets of
	// templates while also providing caching and/or hot-reloading depending on your current environment
	TemplateRenderer struct {
		// templateCache stores a cache of parsed page templates
		templateCache sync.Map

		// funcMap stores the template function map
		funcMap template.FuncMap

		// templatePath stores the complete path to the templates directory
		templatesPath string

		// config stores application configuration
		config *config.Config
	}

	// TemplateParsed is a wrapper around parsed templates which are stored in the TemplateRenderer cache
	TemplateParsed struct {
		// Template is the parsed template
		Template *template.Template

		// build stores the build data used to parse the template
		build *templateBuild
	}

	// templateBuild stores the build data used to parse a template
	templateBuild struct {
		group       string
		key         string
		base        string
		files       []string
		directories []string
	}

	// templateBuilder handles chaining a template parse operation
	templateBuilder struct {
		build    *templateBuild
		renderer *TemplateRenderer
	}
)

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

// Parse creates a template build operation
func (t *TemplateRenderer) Parse() *templateBuilder {
	return &templateBuilder{
		renderer: t,
		build:    &templateBuild{},
	}
}

// GetTemplatesPath gets the complete path to the templates directory
func (t *TemplateRenderer) GetTemplatesPath() string {
	return t.templatesPath
}

// getCacheKey gets a cache key for a given group and ID
func (t *TemplateRenderer) getCacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s:%s", group, key)
	}
	return key
}

// parse parses a set of templates and caches them for quick execution
// If the application environment is set to local, the cache will be bypassed and templates will be
// parsed upon each request so hot-reloading is possible without restarts.
// Also included will be the function map provided by the funcmap package.
func (t *TemplateRenderer) parse(build *templateBuild) (*TemplateParsed, error) {
	var tp *TemplateParsed
	var err error

	switch {
	case build.key == "":
		return nil, errors.New("cannot parse template without key")
	case len(build.files) == 0 && len(build.directories) == 0:
		return nil, errors.New("cannot parse template without files or directories")
	case build.base == "":
		return nil, errors.New("cannot parse template without base")
	}

	// Generate the cache key
	cacheKey := t.getCacheKey(build.group, build.key)

	// Check if the template has not yet been parsed or if the app environment is local, so that
	// templates reflect changes without having the restart the server
	if tp, err = t.Load(build.group, build.key); err != nil || t.config.App.Environment == config.EnvLocal {
		// Initialize the parsed template with the function map
		parsed := template.New(build.base + config.TemplateExt).
			Funcs(t.funcMap)

		// Parse all files provided
		if len(build.files) > 0 {
			for k, v := range build.files {
				build.files[k] = fmt.Sprintf("%s/%s%s", t.templatesPath, v, config.TemplateExt)
			}

			parsed, err = parsed.ParseFiles(build.files...)
			if err != nil {
				return nil, err
			}
		}

		// Parse all templates within the provided directories
		for _, dir := range build.directories {
			dir = fmt.Sprintf("%s/%s/*%s", t.templatesPath, dir, config.TemplateExt)
			parsed, err = parsed.ParseGlob(dir)
			if err != nil {
				return nil, err
			}
		}

		// Store the template so this process only happens once
		tp = &TemplateParsed{
			Template: parsed,
			build:    build,
		}
		t.templateCache.Store(cacheKey, tp)
	}

	return tp, nil
}

// Load loads a template from the cache
func (t *TemplateRenderer) Load(group, key string) (*TemplateParsed, error) {
	load, ok := t.templateCache.Load(t.getCacheKey(group, key))
	if !ok {
		return nil, errors.New("uncached page template requested")
	}

	tmpl, ok := load.(*TemplateParsed)
	if !ok {
		return nil, errors.New("unable to cast cached template")
	}

	return tmpl, nil
}

// Execute executes a template with the given data and provides the output
func (t *TemplateParsed) Execute(data interface{}) (*bytes.Buffer, error) {
	if t.Template == nil {
		return nil, errors.New("cannot execute template: template not initialized")
	}

	buf := new(bytes.Buffer)
	err := t.Template.ExecuteTemplate(buf, t.build.base+config.TemplateExt, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Group sets the cache group for the template being built
func (t *templateBuilder) Group(group string) *templateBuilder {
	t.build.group = group
	return t
}

// Key sets the cache key for the template being built
func (t *templateBuilder) Key(key string) *templateBuilder {
	t.build.key = key
	return t
}

// Base sets the name of the base template to be used during template parsing and execution.
// This should be only the file name without a directory or extension.
func (t *templateBuilder) Base(base string) *templateBuilder {
	t.build.base = base
	return t
}

// Files sets a list of template files to include in the parse.
// This should not include the file extension and the paths should be relative to the templates directory.
func (t *templateBuilder) Files(files ...string) *templateBuilder {
	t.build.files = files
	return t
}

// Directories sets a list of directories that all template files within will be parsed.
// The paths should be relative to the templates directory.
func (t *templateBuilder) Directories(directories ...string) *templateBuilder {
	t.build.directories = directories
	return t
}

// Store parsed the templates and stores them in the cache
func (t *templateBuilder) Store() (*TemplateParsed, error) {
	return t.renderer.parse(t.build)
}

// Execute executes the template with the given data.
// If the template has not already been cached, this will parse and cache the template
func (t *templateBuilder) Execute(data interface{}) (*bytes.Buffer, error) {
	tp, err := t.Store()
	if err != nil {
		return nil, err
	}

	return tp.Execute(data)
}
