package admin

import (
	"embed"
	"strings"
	"text/template"
	"unicode"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

var (
	//go:embed templates
	templateDir embed.FS
)

type Extension struct {
	entc.DefaultExtension
}

func (*Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(
			gen.NewTemplate("admin").
				Funcs(template.FuncMap{
					"fieldName":      fieldName,
					"fieldLabel":     FieldLabel,
					"fieldIsPointer": fieldIsPointer,
				}).
				ParseFS(templateDir, "templates/*tmpl"),
		),
	}
}

func fieldName(name string) string {
	if len(name) == 0 {
		return name
	}

	parts := strings.Split(name, "_")
	for i := 0; i < len(parts); i++ {
		if parts[i] == "id" {
			parts[i] = "ID"
		} else {
			parts[i] = upperFirst(parts[i])
		}
	}

	return strings.Join(parts, "")
}

func FieldLabel(name string) string {
	if len(name) == 0 {
		return name
	}

	parts := strings.Split(name, "_")
	for i := 0; i < len(parts); i++ {
		if parts[i] == "id" {
			parts[i] = "ID"
		}
		if i == 0 {
			parts[i] = upperFirst(parts[i])
		}
	}

	return strings.Join(parts, " ")
}

func fieldIsPointer(f *gen.Field) bool {
	switch {
	case f.Type.Type == field.TypeBool:
		return false
	case f.Optional,
		f.Default,
		f.Sensitive(),
		f.Nillable:
		return true
	}
	return false
}

func upperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	out := []rune(s)
	out[0] = unicode.ToUpper(out[0])
	return string(out)
}
