package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrInvalidSignature = errors.New("invalid signature")
)

func ValidateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}

func ParseRequest(channelSecret string, r *http.Request) (*CallbackRequest, error) {
	defer func() { _ = r.Body.Close() }()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if !ValidateSignature(channelSecret, r.Header.Get("x-line-signature"), body) {
		return nil, ErrInvalidSignature
	}

	var cr CallbackRequest
	err = json.Unmarshal(body, &cr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %w, %s", err, body)
	}
	return &cr, nil
}
