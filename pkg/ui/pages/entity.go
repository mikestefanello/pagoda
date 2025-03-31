package pages

import (
	"fmt"
	"net/http"

	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
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

func AdminEntityAdd(ctx echo.Context, schema *load.Schema) error {
	r := ui.NewRequest(ctx)
	r.Title = fmt.Sprintf("Add %s", "entity") // TODO

	nodes := make(Group, 0)

	for _, f := range schema.Fields {
		switch f.Info.Type {
		case field.TypeString:
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "text",
				Label:     f.Name,
			}))
		case field.TypeTime:
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "datetime",
				Label:     f.Name,
			}))
		case field.TypeBool:
			nodes = append(nodes, P(Textf("%s not supported", f.Name)))
		default:
			nodes = append(nodes, P(Textf("%s not supported", f.Name)))
		}
	}

	for _, e := range schema.Edges {
		if e.Inverse {
			continue
		}
		nodes = append(nodes, InputField(InputFieldParams{
			Name:      e.Name,
			InputType: "number",
			Label:     e.Name,
		}))
	}

	nodes = append(nodes, ControlGroup(
		FormButton("is-primary", "Submit"),
		ButtonLink("/", "is-secondary", "Cancel"),
	), CSRF(r))

	return r.Render(layouts.Admin, Form(Method(http.MethodPost), nodes))
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
