package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/claustra01/calendeye/db"
	"github.com/claustra01/calendeye/google"
)

type TokenRequestBody struct {
	Id       string `json:"id"`
	AuthCode string `json:"code"`
}

func UpdateRefreshToken(w http.ResponseWriter, req *http.Request) {
	defer func() { _ = req.Body.Close() }()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	var reqBody TokenRequestBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	client := google.NewOAuthClient(
		req.Context(),
	)
	token, err := client.GetRefreshToken(reqBody.AuthCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Failed to get token"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = db.UpdateRefreshToken(reqBody.Id, token)
	if err == db.ErrNoRecord {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Refresh token updated"))
	if err != nil {
		log.Println(err)
	}
}
