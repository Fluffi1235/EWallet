FROM golang:1.20.4-alpine AS build_base

WORKDIR /evallet

COPY . .

RUN go mod download

RUN go build ./cmd/api

CMD ["./api"]