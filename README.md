# Сервис назначения ревьюеров для Pull Request’ов

### Сервис, который автоматически назначает ревьюеров на Pull Request’ы (PR), а также позволяет управлять командами и участниками. Взаимодействие происходит через HTTP API.

> * Lang: Go
> * DB: PostgreSQL
> * Migrations: Goose
> * Router: gorilla/mux
> * Host: localhost
> * Port: 8080

## Запуск

```sh
make env
docker compose up
```