FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o name ./cmd/main.go

FROM alpine:3.18

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/name /app/name
COPY /db/postgres/migrations /app/db/postgres/migrations

EXPOSE ${PORT}
EXPOSE ${GRPC_PORT}

CMD ["/app/name"]