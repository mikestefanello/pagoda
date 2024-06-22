package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
	"time"
)

const (
	routeNameCache       = "cache"
	routeNameCacheSubmit = "cache.submit"
)

type (
	Cache struct {
		cache *services.CacheClient
		*services.TemplateRenderer
	}

	cacheForm struct {
		Value string `form:"value"`
		form.Submission
	}
)

func init() {
	Register(new(Cache))
}

func (h *Cache) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.cache = c.Cache
	return nil
}

func (h *Cache) Routes(g *echo.Group) {
	g.GET("/cache", h.Page).Name = routeNameCache
	g.POST("/cache", h.Submit).Name = routeNameCacheSubmit
}

func (h *Cache) Page(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageCache
	p.Title = "Set a cache entry"
	p.Form = form.Get[cacheForm](ctx)

	// Fetch the value from the cache
	value, err := h.cache.Get().
		Key("page_cache_example").
		Fetch(ctx.Request().Context())

	// Store the value in the page so it can be rendered, if found
	switch {
	case err == nil:
		p.Data = value.(string)
	case errors.Is(err, services.ErrCacheMiss):
	default:
		return fail(err, "failed to fetch from cache")
	}

	return h.RenderPage(ctx, p)
}

func (h *Cache) Submit(ctx echo.Context) error {
	var input cacheForm

	if err := form.Submit(ctx, &input); err != nil {
		return err
	}

	// Set the cache
	err := h.cache.Set().
		Key("page_cache_example").
		Data(input.Value).
		Expiration(10 * time.Minute).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to set cache")
	}

	form.Clear(ctx)

	return h.Page(ctx)
}
