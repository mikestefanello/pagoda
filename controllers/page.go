package controllers

import (
	"html/template"
	"net/http"
	"time"

	"goweb/msg"
	"goweb/pager"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

const (
	DefaultItemsPerPage = 20
)

type Page struct {
	AppName    string
	Title      string
	Context    echo.Context
	Reverse    func(name string, params ...interface{}) string
	Path       string
	Data       interface{}
	Layout     string
	Name       string
	IsHome     bool
	IsAuth     bool
	StatusCode int
	Metatags   struct {
		Description string
		Keywords    []string
	}
	Pager pager.Pager
	CSRF  string
	Cache struct {
		Enabled bool
		MaxAge  time.Duration
		Tags    []string
	}
}

func NewPage(c echo.Context) Page {
	p := Page{
		Context:    c,
		Reverse:    c.Echo().Reverse,
		Path:       c.Request().URL.Path,
		StatusCode: http.StatusOK,
		Pager:      pager.NewPager(c, DefaultItemsPerPage),
	}

	p.IsHome = p.Path == "/"

	if csrf := c.Get(middleware.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	return p
}

func (p Page) SetMessage(typ msg.Type, value string) {
	msg.Set(p.Context, typ, value)
}

func (p Page) GetMessages(typ msg.Type) []template.HTML {
	strs := msg.Get(p.Context, typ)
	ret := make([]template.HTML, len(strs))
	for k, v := range strs {
		ret[k] = template.HTML(v)
	}
	return ret
}
