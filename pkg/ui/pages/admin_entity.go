package pages

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
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

func AdminEntityInput(ctx echo.Context, schema admin.EntitySchema, values url.Values) error {
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
		g := make(Group, 0, len(entityList.Columns)+2)
		g = append(g, Th(Text("ID")))
		for _, h := range entityList.Columns {
			g = append(g, Th(Text(h)))
		}
		g = append(g, Th())
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
					ColorInfo,
					r.Path(routenames.AdminEntityEdit(entityTypeName), row.ID),
					"Edit",
				),
				Span(Class("mr-2")),
				ButtonLink(
					ColorError,
					r.Path(routenames.AdminEntityDelete(entityTypeName), row.ID),
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

	return r.Render(layouts.Primary, Group{
		Div(
			Class("form-control mb-2"),
			ButtonLink(
				ColorAccent,
				r.Path(routenames.AdminEntityAdd(entityTypeName)),
				fmt.Sprintf("Add %s", entityTypeName),
			),
		),
		Table(
			Class("table table-zebra mb-2"),
			THead(
				Tr(genHeader()),
			),
			TBody(genRows()),
		),
		Pager(
			entityList.Page,
			r.Path(routenames.AdminEntityAdd(entityTypeName)),
			entityList.HasNextPage,
			"",
		),
	})
}
