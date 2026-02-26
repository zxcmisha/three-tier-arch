compose-db:
	docker compose up -d db

compose-app:
	docker compose up -d app

compose-down:
	docker compose down

migrate-up:
	migrate -path migrations -database "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable" down