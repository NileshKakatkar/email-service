# Email Service

Fiber service that consumes order events from Kafka and sends email notifications through the configured SMTP server.

## Run

```sh
go mod download
go run ./cmd/email
```

Default port: `8084`.
