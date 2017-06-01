package main

import (
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/workers"
	"encoding/json"
	"github.com/kapitol-app/octopus/models"
)

func main() {
	const url = "http://localhost:7979/dummy-data/get-senators"
	f := workers.Fetcher{Url: url}
	data, err := f.Fetch()
	if err != nil {
		logger.Log("Failed to fetch data from:", url, "Error:", err)
		return
	}

	logger.Log("Fetch number of bytes from:", url, "Data:", len(data))

	var resp map[string]interface{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		logger.Log("Failed to decode json with error:", err)
		return
	}

	payload, has := resp["payload"]
	if !has {
		logger.Log("There is no `payload` value in response")
		return
	}

	logger.Log("Payload:", payload)
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		logger.Log("Could not encode payload as json:", err)
		return
	}

	var sens []models.Senator
	err = json.Unmarshal(payloadJson, &sens)
	if err != nil {
		logger.Log("Could not decode senators", err)
		return
	}

	logger.Log("Senators:", sens)
}
