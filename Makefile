.DEFAULT_GOAL := build

postgres:
  docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretbank -d postgres:12-alpine
.PHONY:postgres

createdb: postgres
  docker exec -it postgres12 createdb --username=root --owner=root simple_bank
.PHONY:createdb

dropdb:
  docker exec -it postgres12 dropdb simple_bank
.PHONY:dropdb

migrateup: createdb
  migrate -path migrations/postgresql -database "postgresql://root:secretbank@localhost:5432/simple_bank?sslmode=disable" -verbose up
.PHONY:migrateup

migratedown:
  migrate -path migrations/postgresql -database "postgresql://root:secretbank@localhost:5432/simple_bank?sslmode=disable" -verbose down
.PHONY:migratedown

build: migrateup
  go build main.go
.PHONY:build
  
docs:
  swag init -g adapters/primary/api/rest/router/v1/echo_router.go
.PHONY:docs
