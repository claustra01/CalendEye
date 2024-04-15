package handler

import (
	"context"
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"

	"github.com/claustra01/calendeye/db"
	"github.com/claustra01/calendeye/google"
	"github.com/claustra01/calendeye/linebot"
	"github.com/claustra01/calendeye/openai"
	"github.com/claustra01/calendeye/webhook"
)

var (
	liffUrl = os.Getenv("LIFF_URL")
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
		case webhook.FollowEvent:
			// Register user
			err := db.RegisterUser(e.Source.(webhook.UserSource).UserId)
			if err != nil {
				log.Printf("Cannot register user: %+v\n", err)
				_, err = bot.ReplyMessage(
					&linebot.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []linebot.MessageInterface{
							linebot.NewTextMessage("既にユーザーが登録されているか、予期しないエラーが発生しました。"),
						},
					},
				)
			}
			if err != nil {
				log.Print(err)
			} else {
				log.Println("Sent error reply.")
			}
			// Send reply
			replyText := fmt.Sprintf("友達追加ありがとう!!\nまずはこのリンクからGoogleでログインしてね!!\n%s", liffUrl)
			_, err = bot.ReplyMessage(
				&linebot.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages: []linebot.MessageInterface{
						linebot.NewTextMessage(replyText),
					},
				},
			)
			if err != nil {
				log.Print(err)
			} else {
				log.Println("Sent follow reply.")
			}

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

			case webhook.ImageMessageContent:
				var img image.Image
				var format string

				switch message.ContentProvider.Type {
				case "line":
					img, format, err = bot.FetchLineImage(context.Background(), message.MessageContent.Id)
					if err != nil {
						log.Printf("Failed to fetch image: %+v\n", err)
						_, err = bot.ReplyMessage(
							&linebot.ReplyMessageRequest{
								ReplyToken: e.ReplyToken,
								Messages: []linebot.MessageInterface{
									linebot.NewTextMessage("画像の取得に失敗しました。"),
								},
							},
						)
						if err != nil {
							log.Print(err)
						}
					}
				case "external":
					img, format, err = bot.FetchExternalImage(context.Background(), message.ContentProvider.OriginalContentUrl)
					if err != nil {
						log.Printf("Failed to fetch image: %+v\n", err)
						_, err = bot.ReplyMessage(
							&linebot.ReplyMessageRequest{
								ReplyToken: e.ReplyToken,
								Messages: []linebot.MessageInterface{
									linebot.NewTextMessage("画像の取得に失敗しました。"),
								},
							},
						)
						if err != nil {
							log.Print(err)
						}
					}
				default:
					log.Printf("Unknown contentProvider type: %+v\n", message.ContentProvider.Type)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage("不明な画像です。"),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				refreshToken, err := db.GetRefreshToken(e.Source.(webhook.UserSource).UserId)
				if err != nil {
					log.Printf("Failed to get refresh token: %+v\n", err)
					replyText := fmt.Sprintf("まずはこのリンクからGoogleでログインしてね!!\n%s", liffUrl)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage(replyText),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				googleClient := google.NewOAuthClient(context.Background())
				accessToken, err := googleClient.GetAccessToken(refreshToken)
				if err != nil {
					log.Printf("Failed to get access token: %+v\n", err)
					replyText := fmt.Sprintf("まずはこのリンクからGoogleでログインしてね!!\n%s", liffUrl)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage(replyText),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				gpt, err := openai.NewGpt4Vision(context.Background())
				if err != nil {
					log.Printf("Failed to create GPT-4 client: %+v\n", err)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage("OpenAIのエラーです。"),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				eventJson, err := gpt.Img2Txt(img, format)
				if err != nil {
					log.Printf("Failed to convert image to text: %+v\n", err)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage("画像の解析に失敗しました。"),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				content, err := google.ParseCalendarContent(eventJson)
				if err != nil {
					log.Printf("Failed to parse calendar content: %+v\n", err)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage("イベントが見つかりませんでした。"),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				err = googleClient.RegisterCalenderEvent(content, accessToken)
				if err != nil {
					log.Printf("Failed to register event: %+v\n", err)
					_, err = bot.ReplyMessage(
						&linebot.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []linebot.MessageInterface{
								linebot.NewTextMessage("イベントの登録に失敗しました。"),
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}

				if _, err = bot.ReplyMessage(
					&linebot.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []linebot.MessageInterface{
							linebot.NewTextMessage(fmt.Sprintf("「%s」のイベントを登録したよ!!", content.Summary)),
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent image reply.")
				}

			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}
