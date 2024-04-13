package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/claustra01/calendeye/db"
)

type TokenResponseBody struct {
	RefreshToken string `json:"refresh_token"`
}

func UpdateRefreshToken(w http.ResponseWriter, req *http.Request) {
	defer func() { _ = req.Body.Close() }()

	id := req.URL.Query().Get("id")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var token TokenResponseBody
	err = json.Unmarshal(body, &token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = db.UpdateRefreshToken(id, token.RefreshToken)
	if err == db.ErrNoRecord {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Refresh token updated"))
}
