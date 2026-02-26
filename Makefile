docker-postgres:
	docker run -e POSTGRES_PASSWORD=1234 -p 5432:5432 -v ./out/pgdata:/var/lib/postgresql -d postgres:17

docker-app:
	docker run -p 9091:9091 arch-app

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