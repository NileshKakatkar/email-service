package main

import (
	"log"

	"github.com/nilesh/email-service/internal/email"
)

func main() {
	cfg := email.LoadConfig()
	app := email.NewApp(cfg)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
