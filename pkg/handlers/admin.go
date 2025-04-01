package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/ogent"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/redirect"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

// TODO plugins should create keys dynamically
const entityContextKey = "admin:entity"
const entityIDContextKey = "admin:entity_id"

type Admin struct {
	orm   *ent.Client
	graph *gen.Graph
	ogent *ogent.OgentHandler
}

func init() {
	Register(new(Admin))
}

func (h *Admin) Init(c *services.Container) error {
	h.graph = c.Graph
	h.orm = c.ORM
	h.ogent = ogent.NewOgentHandler(h.orm)
	return nil
}

func (h *Admin) Routes(g *echo.Group) {
	// TODO admin user status middleware
	entities := g.Group("/admin/content")

	for _, p := range h.getEntityPlugins() {
		pg := entities.Group(fmt.Sprintf("/%s", strings.ToLower(p.ID)))
		pg.GET("", h.EntityList(p)).Name = p.RouteNameList()
		pg.POST("", h.EntityList(p)).Name = p.RouteNameListSubmit()
		pg.GET("/add", h.EntityAdd(p)).Name = p.RouteNameAdd()
		pg.POST("/add", h.EntityAddSubmit(p)).Name = p.RouteNameAddSubmit()
		pg.GET("/:id/edit", h.EntityEdit(p), h.entityPluginMiddleware(p)).Name = p.RouteNameEdit()
		pg.POST("/:id/edit", h.EntityEditSubmit(p), h.entityPluginMiddleware(p)).Name = p.RouteNameEditSubmit()
		pg.GET("/:id/delete", h.EntityDelete(p), h.entityPluginMiddleware(p)).Name = p.RouteNameDelete()
		pg.POST("/:id/delete", h.EntityDeleteSubmit(p), h.entityPluginMiddleware(p)).Name = p.RouteNameDeleteSubmit()
	}
}

// TODO, maybe this can be used outside of admin stuff as well?
func (h *Admin) entityPluginMiddleware(plugin AdminEntityPlugin) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid entity ID")
			}

			entity, err := plugin.Load(ctx, h.orm, id)
			switch {
			case err == nil:
				ctx.Set(entityIDContextKey, id)
				ctx.Set(entityContextKey, entity)
				return next(ctx)
			case errors.Is(err, new(ent.NotFoundError)):
				return echo.NewHTTPError(http.StatusNotFound, "entity not found")
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
		}
	}
}

func (h *Admin) EntityList(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		pgr := pager.NewPager(ctx, 25)
		params := pages.AdminEntityListParams{
			Title:       p.LabelPlural,
			Headers:     p.Heading,
			EditRoute:   p.RouteNameEdit(),   // todo remove, pass in plugin
			DeleteRoute: p.RouteNameDelete(), // todo remove, pass in plugin
			Pager:       pgr,
		}
		params.Rows, err = p.List(ctx, h.orm, pgr)
		if err != nil {
			return fail(err, fmt.Sprintf("failed to query %s", p.ID))
		}

		return pages.AdminEntityList(ctx, params)
	}
}

func (h *Admin) EntityAdd(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var schema *load.Schema
		for _, s := range h.graph.Schemas {
			if s.Name == p.ID {
				schema = s
			}
		}
		return pages.AdminEntityAdd(ctx, schema)
	}
}

func (h *Admin) EntityAddSubmit(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var v ogent.CreatePasswordTokenReq // TODO type
		err := ctx.Bind(&v)
		if err != nil {
			return fail(err, fmt.Sprintf("failed to bind create password token request body"))
		}
		fmt.Printf("%+v", v)
		return h.EntityAdd(p)(ctx)
	}
}

func (h *Admin) EntityEdit(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h *Admin) EntityEditSubmit(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h *Admin) EntityDelete(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return pages.AdminEntityDelete(ctx)
	}
}

func (h *Admin) EntityDeleteSubmit(p AdminEntityPlugin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Get(entityIDContextKey).(int)
		if err := p.Delete(ctx, h.orm, id); err != nil {
			return fail(err, fmt.Sprintf("failed to delete %s (ID: %d)", p.ID, id))
		}

		msg.Success(ctx, fmt.Sprintf("Successfully deleted %s.", strings.ToLower(p.Label)))

		return redirect.
			New(ctx).
			Route(p.RouteNameList()).
			Go()
	}
}

// TODO inject orm? move to separate package?
type AdminEntityPlugin struct {
	ID          string
	Label       string
	LabelPlural string
	Heading     []string
	List        func(ctx echo.Context, orm *ent.Client, pgr pager.Pager) ([]pages.AdminEntityListRow, error)
	Load        func(ctx echo.Context, orm *ent.Client, id int) (any, error)
	Delete      func(ctx echo.Context, orm *ent.Client, id int) error
}

func (p *AdminEntityPlugin) RouteNameList() string {
	return fmt.Sprintf("admin:%s_list", p.ID)
}

func (p *AdminEntityPlugin) RouteNameListSubmit() string {
	return fmt.Sprintf("admin:%s_list.submit", p.ID)
}

func (p *AdminEntityPlugin) RouteNameAdd() string {
	return fmt.Sprintf("admin:%s_add", p.ID)
}

func (p *AdminEntityPlugin) RouteNameEdit() string {
	return fmt.Sprintf("admin:%s_edit", p.ID)
}

func (p *AdminEntityPlugin) RouteNameDelete() string {
	return fmt.Sprintf("admin:%s_delete", p.ID)
}

func (p *AdminEntityPlugin) RouteNameAddSubmit() string {
	return fmt.Sprintf("admin:%s_add.submit", p.ID)
}

func (p *AdminEntityPlugin) RouteNameEditSubmit() string {
	return fmt.Sprintf("admin:%s_edit.submit", p.ID)
}

func (p *AdminEntityPlugin) RouteNameDeleteSubmit() string {
	return fmt.Sprintf("admin:%s_delete.submit", p.ID)
}

func (h *Admin) getEntityPlugins() []AdminEntityPlugin {
	return []AdminEntityPlugin{
		{
			ID:          "User",
			Label:       "User",
			LabelPlural: "Users",
			Heading: []string{
				"ID",
				"Name",
				"Email",
				"Created at",
			},
			List: func(ctx echo.Context, client *ent.Client, pgr pager.Pager) ([]pages.AdminEntityListRow, error) {
				users, err := client.User.
					Query().
					Limit(pgr.ItemsPerPage).
					Offset(pgr.GetOffset()).
					Order(user.ByCreatedAt(sql.OrderDesc())).
					All(ctx.Request().Context())

				if err != nil {
					return nil, err
				}

				rows := make([]pages.AdminEntityListRow, 0, len(users))

				for _, u := range users {
					rows = append(rows, pages.AdminEntityListRow{
						ID: u.ID,
						Columns: []string{
							fmt.Sprint(u.ID),
							u.Name,
							u.Email,
							u.CreatedAt.Format(time.RFC822),
						},
					})
				}

				return rows, nil
			},
			Load: func(ctx echo.Context, orm *ent.Client, id int) (any, error) {
				return orm.User.Get(ctx.Request().Context(), id)
			},
			Delete: func(ctx echo.Context, orm *ent.Client, id int) error {
				return orm.User.DeleteOneID(id).Exec(ctx.Request().Context())
			},
		},
		{
			ID:          "PasswordToken",
			Label:       "Password token",
			LabelPlural: "Password tokens",
			Heading: []string{
				"ID",
				"Hash",
				"Created at",
			},
			List: func(ctx echo.Context, client *ent.Client, pgr pager.Pager) ([]pages.AdminEntityListRow, error) {
				tokens, err := client.PasswordToken.
					Query().
					Limit(pgr.ItemsPerPage).
					Offset(pgr.GetOffset()).
					Order(passwordtoken.ByCreatedAt(sql.OrderDesc())).
					All(ctx.Request().Context())

				if err != nil {
					return nil, err
				}

				rows := make([]pages.AdminEntityListRow, 0, len(tokens))

				for _, t := range tokens {
					rows = append(rows, pages.AdminEntityListRow{
						ID: t.ID,
						Columns: []string{
							fmt.Sprint(t.ID),
							t.Hash,
							t.CreatedAt.Format(time.RFC822),
						},
					})
				}

				return rows, nil
			},
			Load: func(ctx echo.Context, orm *ent.Client, id int) (any, error) {
				return orm.PasswordToken.Get(ctx.Request().Context(), id)
			},
			Delete: func(ctx echo.Context, orm *ent.Client, id int) error {
				return orm.PasswordToken.DeleteOneID(id).Exec(ctx.Request().Context())
			},
		},
	}
}
