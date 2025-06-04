package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/spf13/afero"
)

type Build struct {
	files afero.Fs
}

func init() {
	Register(new(Build))
}

func (h *Build) Init(c *services.Container) error {
	return nil
}

func (h *Build) Routes(g *echo.Group) {
	// Serve the build directory
	fs := http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build")))
	g.GET("/build/*", echo.WrapHandler(fs))
}
