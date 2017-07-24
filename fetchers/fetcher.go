package fetchers

import (
	"encoding/json"
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/logger"
	"net/http"
)

func Fetch(url string, response interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-Key", config.C.ApiKeys.ProPublicaCongress)
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error fetching contents from", url, "error:", err)
		return err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		logger.Log("Error: reading contents error:", err)
		return err
	}

	return nil
}
