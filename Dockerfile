FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o users_service ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/users_service .

EXPOSE 9000

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_NAME=messenger_users
ENV GRPC_PORT=9000

CMD ["./users_service"]