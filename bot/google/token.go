package google

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type TokenRequestBody struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
}

type TokenResponseBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func (c *OAuthClient) GetRefreshToken(code string) (string, error) {

	reqBody := TokenRequestBody{
		ClientId:     c.Config.ClientId,
		ClientSecret: c.Config.ClientSecret,
		RedirectUri:  c.Config.RedirectUri,
		GrantType:    "authorization_code",
		Code:         code,
	}
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, c.Config.Endpoint+"/token", bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody TokenResponseBody
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", err
	}

	return respBody.RefreshToken, nil
}
