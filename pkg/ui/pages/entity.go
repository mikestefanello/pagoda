package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Entity(ctx echo.Context) error {
	return ui.NewRequest(ctx).Render(layouts.Admin, Div(Text("abc")))
}

func AdminEntityDelete(ctx echo.Context) error {
	r := ui.NewRequest(ctx)
	form := Form(
		Method(http.MethodPost),
		H2(Text("Are you sure you want to delete this entity?")),
		ControlGroup(
			FormButton("is-link", "Delete"),
			ButtonLink("/", "is-secondary", "Cancel"),
		),
		CSRF(r),
	)

	return r.Render(layouts.Admin, form)
}

type AdminEntityListParams struct {
	Title       string
	Headers     []string
	Rows        []AdminEntityListRow
	EditRoute   string
	DeleteRoute string
	Pager       pager.Pager
}

type AdminEntityListRow struct {
	ID      int
	Columns []string
}

func AdminEntityList(ctx echo.Context, params AdminEntityListParams) error {
	r := ui.NewRequest(ctx)
	r.Title = params.Title

	genHeader := func() Node {
		g := make(Group, 0, len(params.Headers)+2)
		for _, h := range params.Headers {
			g = append(g, Th(Text(h)))
		}
		g = append(g, Th(), Th())
		return g
	}

	genRow := func(row AdminEntityListRow) Node {
		g := make(Group, 0, len(row.Columns)+2)
		for _, h := range row.Columns {
			g = append(g, Td(Text(h)))
		}
		g = append(g,
			Td(A(Href(r.Path(params.EditRoute, row.ID)), Text("Edit"))),
			Td(A(Href(r.Path(params.DeleteRoute, row.ID)), Text("Delete"))),
		)
		return g
	}

	genRows := func() Node {
		g := make(Group, 0, len(params.Rows))
		for _, row := range params.Rows {
			g = append(g, Tr(genRow(row)))
		}
		return g
	}

	return r.Render(layouts.Admin, Table(
		Class("table"),
		THead(
			Tr(genHeader()),
		),
		TBody(genRows()),
	))
}
