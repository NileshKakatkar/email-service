# Email Service

Fiber service that consumes the shared `order.placed` Kafka event contract and sends email notifications through the configured SMTP server.

## Run

```sh
go mod download
go run ./cmd/email
```

Default port: `8084`. Keep `shared/events` beside this service because the event contract is a separate local Go module.
