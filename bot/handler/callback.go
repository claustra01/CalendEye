package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/claustra01/calendeye/linebot"
	"github.com/claustra01/calendeye/webhook"
)

func Callback(w http.ResponseWriter, req *http.Request, bot *linebot.LineBot, channelSecret string) {
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
}
