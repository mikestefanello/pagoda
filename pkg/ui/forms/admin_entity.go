package forms

import (
	"net/http"
	"net/url"

	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AdminEntity(r *ui.Request, schema *load.Schema, values url.Values) Node {
	// TODO inline validation?
	isNew := values == nil
	nodes := make(Group, 0)

	getValue := func(name string) string {
		// Values in the submitted form take precedence.
		if value := r.Context.FormValue(name); value != "" {
			return value
		}

		// Fallback to the entity's values, if being edited.
		if values != nil && len(values[name]) > 0 {
			return values[name][0]
		}

		return ""
	}

	// Attempt to add form elements for all editable entity fields.
	for _, f := range schema.Fields {
		// TODO cardinality?
		if !isNew && f.Immutable {
			continue
		}

		switch f.Info.Type {
		case field.TypeString:
			p := InputFieldParams{
				Name:      f.Name,
				InputType: "text",
				Label:     admin.FieldLabel(f.Name),
				Value:     getValue(f.Name),
			}

			if f.Sensitive {
				p.InputType = "password"
				if !isNew {
					p.Placeholder = "*****"
					p.Help = "SENSITIVE: This field will only be updated if a value is provided."
				}
			}
			nodes = append(nodes, InputField(p))

		case field.TypeTime:
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "datetime-local",
				Label:     admin.FieldLabel(f.Name),
				Value:     getValue(f.Name),
			}))

		case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt64,
			field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64,
			field.TypeFloat32, field.TypeFloat64:
			nodes = append(nodes, InputField(InputFieldParams{
				Name:      f.Name,
				InputType: "number",
				Label:     admin.FieldLabel(f.Name),
				Value:     getValue(f.Name),
			}))

		case field.TypeBool:
			nodes = append(nodes, Checkbox(CheckboxParams{
				Name:    f.Name,
				Label:   admin.FieldLabel(f.Name),
				Checked: getValue(f.Name) == "true",
			}))

		case field.TypeEnum:
			options := make([]Choice, 0, len(f.Enums)+1)
			if f.Optional {
				options = append(options, Choice{
					Label: "-",
					Value: "",
				})
			}
			for _, enum := range f.Enums {
				options = append(options, Choice{
					Label: enum.V,
					Value: enum.V,
				})
			}
			nodes = append(nodes, SelectList(OptionsParams{
				Name:    f.Name,
				Label:   admin.FieldLabel(f.Name),
				Value:   getValue(f.Name),
				Options: options,
			}))

		default:
			nodes = append(nodes, P(Textf("%s not supported", f.Name)))
		}
	}

	return Form(
		Method(http.MethodPost),
		nodes,
		ControlGroup(
			FormButton("btn-primary", "Submit"),
			ButtonLink(
				r.Path(routenames.AdminEntityList(schema.Name)),
				"btn-link",
				"Cancel",
			),
		),
		CSRF(r),
	)
}
