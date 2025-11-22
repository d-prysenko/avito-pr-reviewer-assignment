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