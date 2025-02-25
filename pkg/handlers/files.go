package handlers

import (
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
	"github.com/spf13/afero"
)

type Files struct {
	files afero.Fs
}

func init() {
	Register(new(Files))
}

func (h *Files) Init(c *services.Container) error {
	h.files = c.Files
	return nil
}

func (h *Files) Routes(g *echo.Group) {
	g.GET("/files", h.Page).Name = routenames.Files
	g.POST("/files", h.Submit).Name = routenames.FilesSubmit
}

func (h *Files) Page(ctx echo.Context) error {
	// Compile a list of all uploaded files to be rendered.
	info, err := afero.ReadDir(h.files, "")
	if err != nil {
		return err
	}

	files := make([]*models.File, 0)
	for _, file := range info {
		files = append(files, &models.File{
			Name:     file.Name(),
			Size:     file.Size(),
			Modified: file.ModTime().Format(time.DateTime),
		})
	}

	return pages.UploadFile(ctx, files)
}

func (h *Files) Submit(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		msg.Danger(ctx, "A file is required.")
		return h.Page(ctx)
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
