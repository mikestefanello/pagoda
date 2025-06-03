package pages

import (
	"fmt"
	"net/url"

	"entgo.io/ent/entc/load"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func AdminEntityDelete(ctx echo.Context, entityTypeName string) error {
	r := ui.NewRequest(ctx)
	r.Title = fmt.Sprintf("Delete %s", entityTypeName)

	return r.Render(
		layouts.Primary,
		forms.AdminEntityDelete(r, entityTypeName),
	)
}

func AdminEntityInput(ctx echo.Context, schema *load.Schema, values url.Values) error {
	r := ui.NewRequest(ctx)
	if values == nil {
		r.Title = fmt.Sprintf("Add %s", schema.Name)
	} else {
		r.Title = fmt.Sprintf("Edit %s", schema.Name)
	}

	return r.Render(
		layouts.Primary,
		forms.AdminEntity(r, schema, values),
	)
}

func AdminEntityList(
	ctx echo.Context,
	entityTypeName string,
	entityList *admin.EntityList,
) error {
	r := ui.NewRequest(ctx)
	r.Title = entityTypeName

	genHeader := func() Node {
		g := make(Group, 0, len(entityList.Columns)+3)
		g = append(g, Th(Text("ID")))
		for _, h := range entityList.Columns {
			g = append(g, Th(Text(h)))
		}
		g = append(g, Th(), Th())
		return g
	}

	genRow := func(row admin.EntityValues) Node {
		g := make(Group, 0, len(row.Values)+3)
		g = append(g, Th(Text(fmt.Sprint(row.ID))))
		for _, h := range row.Values {
			g = append(g, Td(Text(h)))
		}
		g = append(g,
			Td(
				ButtonLink(
					r.Path(routenames.AdminEntityEdit(entityTypeName), row.ID),
					"btn-primary",
					"Edit",
				),
			),
			Td(
				ButtonLink(r.Path(routenames.AdminEntityDelete(entityTypeName), row.ID),
					"btn-error",
					"Delete",
				),
			),
		)
		return g
	}

	genRows := func() Node {
		g := make(Group, 0, len(entityList.Entities))
		for _, row := range entityList.Entities {
			g = append(g, Tr(genRow(row)))
		}
		return g
	}

	pagedHref := func(page int) string {
		return fmt.Sprintf("%s?%s=%d",
			r.Path(routenames.AdminEntityList(entityTypeName)),
			pager.QueryKey,
			page,
		)
	}

	return r.Render(layouts.Primary, Group{
		ButtonLink(
			r.Path(routenames.AdminEntityAdd(entityTypeName)),
			"btn-primary",
			fmt.Sprintf("Add %s", entityTypeName),
		),
		Table(
			Class("table"),
			THead(
				Tr(genHeader()),
			),
			TBody(genRows()),
		),
		Nav(
			Class("pagination"),
			A(
				Classes{
					"pagination-previous": true,
					"is-disabled":         entityList.Page == 1,
				},
				If(entityList.Page != 1, Href(pagedHref(entityList.Page-1))),
				Text("Previous page"),
			),
			A(
				Classes{
					"pagination-previous": true,
					"is-disabled":         !entityList.HasNextPage,
				},
				If(entityList.HasNextPage, Href(pagedHref(entityList.Page+1))),
				Text("Next page"),
			),
		),
	})
}
