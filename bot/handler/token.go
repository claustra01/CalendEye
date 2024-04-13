package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/claustra01/calendeye/db"
)

type TokenResponseBody struct {
	RefreshToken string `json:"refresh_token"`
}

func UpdateRefreshToken(w http.ResponseWriter, req *http.Request) {
	defer func() { _ = req.Body.Close() }()

	id := req.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("id is required"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	var token TokenResponseBody
	err = json.Unmarshal(body, &token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = db.UpdateRefreshToken(id, token.RefreshToken)
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
