package handlers

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui"
)

type Cache struct {
	cache *services.CacheClient
}

func init() {
	Register(new(Cache))
}

func (h *Cache) Init(c *services.Container) error {
	h.cache = c.Cache
	return nil
}

func (h *Cache) Routes(g *echo.Group) {
	g.GET("/cache", h.Page).Name = routenames.Cache
	g.POST("/cache", h.Submit).Name = routenames.CacheSubmit
}

func (h *Cache) Page(ctx echo.Context) error {
	f := form.Get[ui.CacheForm](ctx)

	// Fetch the value from the cache.
	value, err := h.cache.
		Get().
		Key("page_cache_example").
		Fetch(ctx.Request().Context())

	// Store the value in the form, so it can be rendered, if found.
	switch {
	case err == nil:
		f.CurrentValue = value.(string)
	case errors.Is(err, services.ErrCacheMiss):
	default:
		return fail(err, "failed to fetch from cache")
	}

	return ui.UpdateCache(ctx, f)
}

func (h *Cache) Submit(ctx echo.Context) error {
	var input ui.CacheForm

	if err := form.Submit(ctx, &input); err != nil {
		return err
	}

	// Set the cache.
	err := h.cache.
		Set().
		Key("page_cache_example").
		Data(input.Value).
		Expiration(30 * time.Minute).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to set cache")
	}

	form.Clear(ctx)

	return h.Page(ctx)
}
