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
					"fieldLabel":     fieldLabel,
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

func fieldLabel(name string) string {
	if len(name) == 0 {
		return name
	}

	out := strings.ReplaceAll(name, "_", " ")
	return upperFirst(out)
}

func fieldIsPointer(f *gen.Field) bool {
	switch {
	case f.Type.Type == field.TypeBool:
		return false
	case f.Optional,
		f.Default:
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

/*
TODO:
1) How to handle fields like password that need to be transformed or omitted, etc?
2) Should we use the HTML datetime format and string fields rather than time.Time?
*/
