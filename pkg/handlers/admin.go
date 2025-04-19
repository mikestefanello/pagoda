package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/redirect"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

// TODO plugins should create keys dynamically
const entityContextKey = "admin:entity"
const entityIDContextKey = "admin:entity_id"

type Admin struct {
	orm   *ent.Client
	graph *gen.Graph
	admin *admin.Handler
}

func init() {
	Register(new(Admin))
}

func (h *Admin) Init(c *services.Container) error {
	h.graph = c.Graph
	h.orm = c.ORM
	h.admin = admin.NewHandler(h.orm, admin.HandlerConfig{
		ItemsPerPage: 25,
		PageQueryKey: pager.QueryKey,
		TimeFormat:   time.DateTime,
	})
	return nil
}

func (h *Admin) Routes(g *echo.Group) {
	// TODO admin user status middleware
	entities := g.Group("/admin/content")

	for _, n := range h.graph.Nodes {
		ng := entities.Group(fmt.Sprintf("/%s", strings.ToLower(n.Name)))
		ng.GET("", h.EntityList(n)).
			Name = routenames.AdminEntityList(n.Name)
		ng.POST("", h.EntityList(n)).
			Name = routenames.AdminEntityListSubmit(n.Name)
		ng.GET("/add", h.EntityAdd(n)).
			Name = routenames.AdminEntityAdd(n.Name)
		ng.POST("/add", h.EntityAddSubmit(n)).
			Name = routenames.AdminEntityAddSubmit(n.Name)
		ng.GET("/:id/edit", h.EntityEdit(n), h.middlewareEntityLoad(n)).
			Name = routenames.AdminEntityEdit(n.Name)
		ng.POST("/:id/edit", h.EntityEditSubmit(n), h.middlewareEntityLoad(n)).
			Name = routenames.AdminEntityEditSubmit(n.Name)
		ng.GET("/:id/delete", h.EntityDelete(n), h.middlewareEntityLoad(n)).
			Name = routenames.AdminEntityDelete(n.Name)
		ng.POST("/:id/delete", h.EntityDeleteSubmit(n), h.middlewareEntityLoad(n)).
			Name = routenames.AdminEntityDeleteSubmit(n.Name)
	}
}

func (h *Admin) middlewareEntityLoad(n *gen.Type) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid entity ID")
			}

			entity, err := h.admin.Get(ctx, n.Name, id)
			switch {
			case err == nil:
				ctx.Set(entityIDContextKey, id)
				ctx.Set(entityContextKey, map[string][]string(entity))
				return next(ctx)
			case ent.IsNotFound(err):
				return echo.NewHTTPError(http.StatusNotFound, "entity not found")
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
		}
	}
}

func (h *Admin) EntityList(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		list, err := h.admin.List(ctx, n.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return pages.AdminEntityList(ctx, pages.AdminEntityListParams{
			EntityType: n,
			EntityList: list,
			Pager:      pager.NewPager(ctx, h.admin.Config.ItemsPerPage),
		})
	}
}

func (h *Admin) EntityAdd(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return pages.AdminEntityForm(ctx, true, h.getEntitySchema(n), nil)
	}
}

func (h *Admin) EntityAddSubmit(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := h.admin.Create(ctx, n.Name)
		if err != nil {
			msg.Danger(ctx, err.Error())
			return h.EntityAdd(n)(ctx)
		}

		msg.Success(ctx, fmt.Sprintf("Successfully added %s.", n.Name))

		return redirect.
			New(ctx).
			Route(routenames.AdminEntityList(n.Name)).
			StatusCode(http.StatusFound).
			Go()
	}
}

func (h *Admin) EntityEdit(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		v := ctx.Get(entityContextKey).(map[string][]string)
		return pages.AdminEntityForm(ctx, false, h.getEntitySchema(n), v)
	}
}

func (h *Admin) EntityEditSubmit(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Get(entityIDContextKey).(int)
		err := h.admin.Update(ctx, n.Name, id)
		if err != nil {
			msg.Danger(ctx, err.Error())
			return h.EntityEdit(n)(ctx)
		}

		msg.Success(ctx, fmt.Sprintf("Updated %s.", n.Name))

		return redirect.
			New(ctx).
			Route(routenames.AdminEntityList(n.Name)).
			StatusCode(http.StatusFound).
			Go()
	}
}

func (h *Admin) EntityDelete(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return pages.AdminEntityDelete(ctx, n.Name)
	}
}

func (h *Admin) EntityDeleteSubmit(n *gen.Type) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Get(entityIDContextKey).(int)
		if err := h.admin.Delete(ctx, n.Name, id); err != nil {
			msg.Danger(ctx, err.Error())
			return h.EntityDelete(n)(ctx)
		}

		msg.Success(ctx, fmt.Sprintf("Successfully deleted %s (ID %d).", n.Name, id))

		return redirect.
			New(ctx).
			Route(routenames.AdminEntityList(n.Name)).
			StatusCode(http.StatusFound).
			Go()
	}
}

func (h *Admin) getEntitySchema(n *gen.Type) *load.Schema {
	for _, s := range h.graph.Schemas {
		if s.Name == n.Name {
			return s
		}
	}
	return nil
}
