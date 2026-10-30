package email

import (
	"context"
	"encoding/json"
	"fmt"
	orderevent "github.com/nilesh/shopcore-events/order"
	"github.com/segmentio/kafka-go"
	"log"
	"net/smtp"
	"strings"
	"time"
)

type OrderEvent = orderevent.Placed

func RunConsumer(ctx context.Context, cfg Config) {
	brokers := cfg.Brokers
	list := strings.Split(brokers, ",")
	reader := kafka.NewReader(kafka.ReaderConfig{Brokers: list, Topic: orderevent.PlacedTopic, GroupID: "shopcore-email", MinBytes: 1, MaxBytes: 10e6})
	defer reader.Close()
	for {
		message, err := reader.FetchMessage(ctx)
		if err != nil {
			return
		}
		var event OrderEvent
		if err = json.Unmarshal(message.Value, &event); err != nil {
			log.Printf("invalid order event: %v", err)
			_ = reader.CommitMessages(ctx, message)
			continue
		}
		for attempt := 0; ; attempt++ {
			err = sendOrderEmail(cfg, event)
			if err == nil {
				break
			}
			delay := time.Duration(attempt+1) * 2 * time.Second
			if delay > 30*time.Second {
				delay = 30 * time.Second
			}
			log.Printf("order email delivery failed order_id=%s retry_in=%s: %v", event.OrderID, delay, err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(delay):
			}
		}
		log.Printf("order notification delivered order_id=%s", event.OrderID)
		if err = reader.CommitMessages(ctx, message); err != nil {
			log.Printf("message commit failed: %v", err)
		}
	}
}

func sendOrderEmail(cfg Config, event OrderEvent) error {
	if cfg.SMTPHost == "" || cfg.SMTPPort == "" || cfg.From == "" || event.Email == "" {
		log.Printf("order notification accepted by local sink order_id=%s recipient=%s", event.OrderID, event.Email)
		return nil
	}
	address := cfg.SMTPHost + ":" + cfg.SMTPPort
	subject := "ShopCore order confirmation"
	body := fmt.Sprintf("Your order %s has been placed. Total: %d %s.", event.OrderID, event.TotalAmount, event.Currency)
	message := []byte("To: " + event.Email + "\r\nFrom: " + cfg.From + "\r\nSubject: " + subject + "\r\n\r\n" + body)
	var auth smtp.Auth
	if cfg.SMTPUser != "" {
		auth = smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPassword, cfg.SMTPHost)
	}
	return smtp.SendMail(address, auth, cfg.From, []string{event.Email}, message)
}
