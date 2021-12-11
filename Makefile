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

.PHONY: up
up:
	docker-compose up -d
	sleep 3

.PHONY: run
run:
	clear
	go run main.go