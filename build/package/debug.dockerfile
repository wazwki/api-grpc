FROM golang:1.23.2

WORKDIR /app

EXPOSE ${PORT}
EXPOSE ${GRPC_PORT}
EXPOSE 8081

CMD ["go", "run", "cmd/name/main.go"]