//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

var returnAllErrors = gen.MustParse(gen.NewTemplate("").Parse(`
{{ define "ogent/ogent/helper/error" }}{{/* gotype: entgo.io/ent/entc/gen.typeScope */}}
	{{- $pkg := base $.Type.Config.Package }}
	if err != nil {
		{{- with $.Scope.Tx }}
			if rErr := {{ . }}.Rollback(); rErr != nil {
				return nil, fmt.Errorf("%w: %v", err, rErr)
			}
		{{- end }}
		// Let the server handle the error.
		return nil, err
	}
{{ end }}
`))

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec, ogent.Templates(returnAllErrors))
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
