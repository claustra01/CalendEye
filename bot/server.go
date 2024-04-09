package main

import (
	"log"
	"net/http"
	"os"

	"github.com/claustra01/calendeye/db"
	"github.com/claustra01/calendeye/handler"
	"github.com/claustra01/calendeye/linebot"
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

	// Connect to database
	err := db.DB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	// Auto migrate
	err = db.DB.Migrate()
	if err != nil {
		log.Fatal(err)
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
		handler.Callback(w, req, bot, channelSecret)
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
