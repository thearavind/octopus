package fetchers

import (
	"bytes"
	"net/url"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/endpoints"
	"github.com/kapitol-app/octopus/logger"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func FetchTwitterBearerToken() string {
	client := &http.Client{}
	token := tokenResponse{}
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", endpoints.GetEndpoint(endpoints.TwitterAccessToken),
		bytes.NewBufferString(data.Encode()))
	req.Header.Add("Authorization", "Basic "+ base64.StdEncoding.EncodeToString([]byte(
		config.C.TwitterKeys.ConsumerKey+ ":"+ config.C.TwitterKeys.ConsumerSecret)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error fetching twitter token", err)
		return ""
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&token)

	if err != nil {
		logger.Log("Error: reading twitter token response contents:", err)
		return ""
	}
	return token.AccessToken
}

func FetchTweets(twitterProfile string, bearer string, response interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoints.GetEndpoint(endpoints.TwitterSearch)+"from%40"+twitterProfile, nil)
	req.Header.Add("Authorization", "Bearer "+bearer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error fetching member", twitterProfile, "tweet result", err)
		return err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		logger.Log("Error: reading twitter search response contents:", err)
		return err
	}
	return nil
}
