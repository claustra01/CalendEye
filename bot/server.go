package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
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
	bot, err := messaging_api.NewMessagingApiAPI(channelToken)
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
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
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
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								messaging_api.TextMessage{
									Text: message.Text,
								},
							},
						},
					); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent text reply.")
					}
				case webhook.StickerMessageContent:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerId, message.StickerResourceType)
					if _, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								messaging_api.TextMessage{
									Text: replyMessage,
								},
							},
						}); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent sticker reply.")
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
	fmt.Println("http://localhost:" + port + "/")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
