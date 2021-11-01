package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	clickup "github.com/flavioltonon/clickup-go-client"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	client, err := clickup.NewClient(
		&http.Client{
			Timeout: 10 * time.Second,
		},
		clickup.SetPersonalTokenAuthorization(os.Getenv("PERSONAL_TOKEN")),
	)
	if err != nil {
		log.Fatal(err)
	}

	webhooks, err := client.Teams.ListWebhooks("30904221")
	if err != nil {
		log.Fatal(err)
	}

	for _, webhook := range webhooks {
		fmt.Printf("webhook.GetID(): %v\n", webhook.GetID())
		fmt.Printf("webhook.GetEndpoint(): %v\n", webhook.GetEndpoint())
	}
}
