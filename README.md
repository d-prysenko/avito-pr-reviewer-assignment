# Сервис назначения ревьюеров для Pull Request’ов

### Сервис, который автоматически назначает ревьюеров на Pull Request’ы (PR), а также позволяет управлять командами и участниками. Взаимодействие происходит через HTTP API.

> * Lang: Go
> * DB: PostgreSQL
> * Migrations: Goose
> * Router: gorilla/mux
> * Validator: go-playground/validator
> * Host: localhost
> * Port: 8080

## Запуск

```sh
make env
make up # или docker compose up -d
```
Приложение собирается и работает внутри докера, реализовано api из ТЗ (docs/openapi.yml), _хотел добавить еще несколько эндпоинтов от себя, но не успел, так же как и добавить multistage сборку, пуш образов в репозиторий и много чего другого_. 
```sh
# Линтер работает локально. vet + fmt + staticcheck
make lint
```
## Тесты
```sh
make test
```
Для тестов используется /docker-compose.testing.yaml, /tests/.env.testing, через докер поднимается отдельная база данных и прогоняются e2e тесты на основные кейсы. _TODO: замокать всё, покрыть юнитами_