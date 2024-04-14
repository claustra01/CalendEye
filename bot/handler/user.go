package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/claustra01/calendeye/db"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
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

	user, err := db.GetUser(id)
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

	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(jsonUser))
	if err != nil {
		log.Println(err)
	}
}
