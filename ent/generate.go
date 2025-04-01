package ent

//go:generate go run -mod=mod entc.go
//go:generate sed -i -e "s/json:\"/form:\"/g" ogent/oas_schemas_gen.go
