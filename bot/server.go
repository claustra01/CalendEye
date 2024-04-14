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

// CORS middleware
func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		}
		next.ServeHTTP(w, r)
	})
}

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

	// Setup HTTP Server for get/update user information
	http.HandleFunc("/user", handler.GetUser)
	http.HandleFunc("/token", handler.UpdateRefreshToken)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Server started!")
	err = http.ListenAndServe(":"+port, CORSHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
