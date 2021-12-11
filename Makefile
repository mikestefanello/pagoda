.PHONY: pg
pg:
	 psql postgresql://admin:admin@localhost:5432/app

.PHONY: ent-gen
ent-gen:
	go generate ./ent

.PHONY: ent-new
ent-new:
	go run entgo.io/ent/cmd/ent init $(name)

.PHONY: ent-install
ent-install:
	go get -d entgo.io/ent/cmd/ent