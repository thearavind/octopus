package workers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/logger"
)

const base = "https://api.propublica.org/congress/v1/"

//Chamber - Chamber of the congress
type Chamber string

const (
	House Chamber = "house"
	Senate Chamber = "senate"
)

//Type - Type of the api query
type Type string

const (
	//Member - member type
	Member Type = "members"
)

//Fetch - Fetches the data from the propublica api
func Fetch(congress int, chamber Chamber, ty Type, end string, response interface{}) error {
	query := fmt.Sprintf("%s%d/%s/%s%s", base, congress, chamber, ty, end)
	client := &http.Client{}
	req, err := http.NewRequest("GET", query, nil)
	req.Header.Add("X-API-Key", config.C.ApiKeys.ProPublicaCongress)
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error fetching contents from", query, "error:", err)
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
