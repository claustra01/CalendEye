package google

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type RefreshTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
}

type AccessTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func (c *OAuthClient) GetRefreshToken(code string) (string, error) {

	reqBody := RefreshTokenRequest{
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

	// TODO: Clone context
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.Config.Endpoint, bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody TokenResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", err
	}

	return respBody.RefreshToken, nil
}

func (c *OAuthClient) GetAccessToken(refreshToken string) (string, error) {

	reqBody := AccessTokenRequest{
		ClientId:     c.Config.ClientId,
		ClientSecret: c.Config.ClientSecret,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// TODO: Clone context
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.Config.Endpoint, bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody TokenResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", err
	}

	return respBody.AccessToken, nil
}
