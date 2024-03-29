package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/claustra01/calendeye/linebot"
	"github.com/claustra01/calendeye/webhook"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if os.Getenv("GOENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Get channel secret and channel token from environment variables
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")

	// Initialize LINE bot
	bot, err := linebot.NewBot(channelToken)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		log.Println("/callback called...")

		cb, err := webhook.ParseRequest(channelSecret, req)
		if err != nil {
			log.Printf("Cannot parse request: %+v\n", err)
			if errors.Is(err, webhook.ErrInvalidSignature) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		log.Println("Handling events...")
		for _, event := range cb.Events {
			log.Printf("/callback called%+v...\n", event)

			switch e := event.(type) {
			case webhook.MessageEvent:
				switch message := e.Message.(type) {
				case webhook.TextMessageContent:
					if _, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage(message.Text),
							},
						},
					); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent text reply.")
					}
				default:
					log.Printf("Unsupported message content: %T\n", e.Message)
				}
			default:
				log.Printf("Unsupported message: %T\n", event)
			}
		}
	})

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Server started!")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
