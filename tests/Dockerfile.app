# syntax=docker/dockerfile:1

FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /revass ./cmd/main.go

RUN cp tests/.env.testing .env

CMD ["/revass"]
