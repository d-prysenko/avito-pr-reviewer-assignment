.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:vet
vet: fmt
	go vet ./...

.PHONY:lint
lint: vet
	staticcheck ./...

env:
	cp .env.dist .env

up:
	docker compose up -d --build

.PHONY:test
test:
	docker compose --env-file tests/.env.testing -f docker-compose.testing.yaml up --build --abort-on-container-exit