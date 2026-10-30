package email

import "os"

type Config struct {
	Port         string
	Brokers      string
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	From         string
}

func LoadConfig() Config {
	port := os.Getenv("EMAIL_SERVICE_PORT")
	if port == "" {
		port = "8084"
	}
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}
	return Config{Port: port, Brokers: brokers, SMTPHost: os.Getenv("SMTP_HOST"), SMTPPort: os.Getenv("SMTP_PORT"), SMTPUser: os.Getenv("SMTP_USER"), SMTPPassword: os.Getenv("SMTP_PASSWORD"), From: os.Getenv("SMTP_FROM")}
}
