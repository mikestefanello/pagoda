package pages

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AdminEntityDelete(ctx echo.Context, entityTypeName string) error {
	r := ui.NewRequest(ctx)
	r.Title = fmt.Sprintf("Delete %s", entityTypeName)

	form := Form(
		Method(http.MethodPost),
		H2(Textf("Are you sure you want to delete this %s?", entityTypeName)),
		ControlGroup(
			FormButton("is-link", "Delete"),
			ButtonLink(r.Path(routenames.AdminEntityList(entityTypeName)), "is-secondary", "Cancel"),
		),
		CSRF(r),
	)

	return r.Render(layouts.Admin, form)
}

func AdminEntityForm(ctx echo.Context, isNew bool, schema *load.Schema, values url.Values) error {
	r := ui.NewRequest(ctx)
	if isNew {
		r.Title = fmt.Sprintf("Add %s", schema.Name)
	} else {
		r.Title = fmt.Sprintf("Edit %s", schema.Name)
	}

	nodes := make(Group, 0)

	label := func(name string) string {
		if len(name) == 0 {
			return name
		}
		text := []rune(strings.ReplaceAll(name, "_", " "))
		text[0] = unicode.ToUpper(text[0])
		return string(text)
	}

	getValue := func(name string) string {
		if value := ctx.FormValue(name); value != "" {
			return value
		}

		if values != nil && len(values[name]) > 0 {
			return values[name][0] // TODO cardinality
		}

		return ""
	}

	for _, f := range schema.Fields {
		// TODO cardinality?
		if !isNew && f.Immutable {
			continue
		}
		// TODO sensitive edits
		switch f.Info.Type {
		case field.TypeString:
			inputType := "text"
			if f.Sensitive {
				inputType = "password"
			}
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: inputType,
				Label:     label(f.Name),
				Value:     getValue(f.Name),
			}))
		case field.TypeTime:
			// todo make this easier
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "text",
				Label:     label(f.Name),
				Help:      fmt.Sprintf("Use the following format: %s", time.Now().Format(time.RFC3339)),
				Value:     getValue(f.Name),
			}))
		case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt64,
			field.TypeFloat32, field.TypeFloat64:
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "number",
				Label:     label(f.Name),
				Value:     getValue(f.Name),
			}))
		case field.TypeBool:
			nodes = append(nodes, Checkbox(CheckboxParams{
				Name:    f.Name,
				Label:   label(f.Name),
				Checked: getValue(f.Name) == "true",
			}))
		case field.TypeEnum:
			// TODO
			nodes = append(nodes, P(Textf("%s not supported", f.Name)))
		default:
			nodes = append(nodes, P(Textf("%s not supported", f.Name)))
		}
	}

	nodes = append(nodes, ControlGroup(
		FormButton("is-primary", "Submit"),
		ButtonLink(r.Path(routenames.AdminEntityList(schema.Name)), "is-secondary", "Cancel"),
	), CSRF(r))

	return r.Render(layouts.Admin, Form(
		Method(http.MethodPost),
		nodes,
	))
}

type AdminEntityListParams struct {
	EntityType *gen.Type
	EntityList *admin.EntityList
	Pager      pager.Pager
}

func AdminEntityList(ctx echo.Context, params AdminEntityListParams) error {
	r := ui.NewRequest(ctx)
	r.Title = params.EntityType.Name

	genHeader := func() Node {
		g := make(Group, 0, len(params.EntityList.Columns)+3)
		g = append(g, Th(Text("ID")))
		for _, h := range params.EntityList.Columns {
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
					r.Path(routenames.AdminEntityEdit(params.EntityType.Name), row.ID),
					"is-link",
					"Edit",
				),
			),
			Td(
				ButtonLink(r.Path(routenames.AdminEntityDelete(params.EntityType.Name), row.ID),
					"is-danger",
					"Delete",
				),
			),
		)
		return g
	}

	genRows := func() Node {
		g := make(Group, 0, len(params.EntityList.Entities))
		for _, row := range params.EntityList.Entities {
			g = append(g, Tr(genRow(row)))
		}
		return g
	}

	// TODO pager
	return r.Render(layouts.Admin, Group{
		ButtonLink(
			r.Path(routenames.AdminEntityAdd(params.EntityType.Name)),
			"is-link",
			fmt.Sprintf("Add %s", params.EntityType.Name),
		),
		Table(
			Class("table"),
			THead(
				Tr(genHeader()),
			),
			TBody(genRows()),
		),
	})
}
