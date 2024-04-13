package handler

import (
	"encoding/json"
	"net/http"

	"github.com/claustra01/calendeye/db"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	defer func() { _ = req.Body.Close() }()

	id := req.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is required"))
		return
	}

	user, err := db.GetUser(id)
	if err == db.ErrNoRecord {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonUser))
}
