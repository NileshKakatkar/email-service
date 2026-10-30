module github.com/nilesh/email-service

go 1.23

replace github.com/nilesh/shopcore-events => ../shared/events

require (
	github.com/nilesh/shopcore-events v0.0.0

	github.com/gofiber/fiber/v2 v2.52.6
	github.com/segmentio/kafka-go v0.4.47
)
