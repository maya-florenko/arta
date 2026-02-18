default: run

run:
	go run cmd/app/main.go

up:
	docker compose up -d --build
