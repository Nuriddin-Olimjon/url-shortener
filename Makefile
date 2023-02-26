run:
	docker compose up -d --build

stop:
	docker compose down

restart:
	docker compose down
	docker compose up -d --build

logs:
	docker compose logs $(name) -f

exec:
	docker compose exec -it $(name) bash

sqlc:
	sqlc generate

swag:
	swag init -g ./main.go -o ./docs/api/swagger

migrate_create:
	migrate create -ext sql -dir ./internal/repository/migrations -seq ${FILE_NAME}

migrateup:
	migrate -path ./internal/repository/migrations -database ${PSQL_URL} -verbose up $(N)

migratedown:
	migrate -path ./internal/repository/migrations -database ${PSQL_URL} -verbose down $(N)

.PHONY: run stop restart logs exec sqlc migrate_create migrateup migratedown
