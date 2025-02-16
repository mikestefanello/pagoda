package handlers

import (
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
	"github.com/spf13/afero"
)

const (
	routeNameFiles       = "files"
	routeNameFilesSubmit = "files.submit"
)

type (
	Files struct {
		files afero.Fs
		*services.TemplateRenderer
	}

	File struct {
		Name     string
		Size     int64
		Modified string
	}
)

func init() {
	Register(new(Files))
}

func (h *Files) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.files = c.Files
	return nil
}

func (h *Files) Routes(g *echo.Group) {
	g.GET("/files", h.Page).Name = routeNameFiles
	g.POST("/files", h.Submit).Name = routeNameFilesSubmit
}

func (h *Files) Page(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageFiles
	p.Title = "Upload a file"

	// Send a list of all uploaded files to the template to be rendered.
	info, err := afero.ReadDir(h.files, "")
	if err != nil {
		return err
	}

	files := make([]File, 0)
	for _, file := range info {
		files = append(files, File{
			Name:     file.Name(),
			Size:     file.Size(),
			Modified: file.ModTime().Format(time.DateTime),
		})
	}

	p.Data = files

	return h.RenderPage(ctx, p)
}

func (h *Files) Submit(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := h.files.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	msg.Success(ctx, fmt.Sprintf("%s was uploaded successfully.", file.Filename))

	return h.Page(ctx)
}
