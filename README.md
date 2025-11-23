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
make env # или cp .env.dist .env
make up # или docker compose up -d
```
Приложение собирается и работает внутри докера, реализовано api из ТЗ (docs/openapi.yml), _чисто технически это, конечно, запуск не одной командой docker compose up, а двумя командами, но .env в репозиторий закидывать это какое-то грубиянство. Хотел добавить еще несколько эндпоинтов от себя, но не успел, так же как и добавить multistage сборку, пуш образов в репозиторий и много чего другого. А еще тут интерфейсы расположены не по месту использования_. 
```sh
# Линтер работает локально. vet + fmt + staticcheck
make lint
```
## Тесты
```sh
make test
```
Для тестов используется /docker-compose.testing.yaml, /tests/.env.testing, в одном контейнере поднимается приложение, в другом отдельная база данных, в третьем прогоняются e2e тесты на основные кейсы. _TODO: замокать всё, покрыть юнитами_

### Парочка простеньких k6 нагрузочных тестов

Получение команды
![Получение команды](https://github.com/d-prysenko/avito-pr-reviewer-assignment/blob/main/docs/res/team-get.jpg?raw=true)

Добавление команды
![Добавление команды](https://github.com/d-prysenko/avito-pr-reviewer-assignment/blob/main/docs/res/team-add.jpg?raw=true)