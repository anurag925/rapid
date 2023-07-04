include .env

# DB migrations
migrations_location="./pkg/models/migrations"
go_migrate=migrate -verbose
migrater=${go_migrate} -path ${migrations_location} -database mysql://${DB_URL}

all:
	go run cmd/server/main.go 

build:
	go build cmd/server/main.go
run:
	./main
migrate:
	go run cmd/migrate/main.go 
script:
	go run cmd/script/main.go 

# Migration
generate_migration:
	${go_migrate} create -ext sql -tz utc -dir ${migrations_location} ${name}

# migrate_up: 
#     ${migrater} up
#   migrate -verbose -path "./pkg/models/migrations" -database "mysql://root@tcp(127.0.0.1:3306)/rapid?parseTime=true" up

# migrate_up_to:
#     ${migrate} up ${version}

# migrate_down_all:
#     ${migrate} down

# migrate_down:
#     ${migrate} down 1

# migrate_down_by:
#     ${migrate} down ${versions}

# migrate_version:
#     ${migrate} version

# migrate_drop:
#     ${migrate} drop

# migrate_force:
#     ${migrate} force ${version}