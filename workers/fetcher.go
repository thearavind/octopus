package workers

import (
	"encoding/json"
	"fmt"
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/endpoints"
	"github.com/kapitol-app/octopus/logger"
	"net/http"
)

//Chamber - Chamber of the congress
type Chamber string

const (
	House  Chamber = "house"
	Senate Chamber = "senate"
)

//Type - Type of the api query
type FetchType string

const (
	//Member - member type
	Members FetchType = "members"
)

//Fetch - Fetches the data from the propublica api
func PropublicaMembersFetch(
	congress int,
	chamber Chamber,
	response interface{},
) error {
	url := fmt.Sprintf(
		"%s%d/%s/%s",
		endpoints.GetEndpoint(endpoints.Propublica),
		congress,
		chamber,
		"members.json",
	)

	return fetch(url, response)
}

func PropublicaMemberFetch(
	memberUrl string,
	response interface{},
) error {

	return fetch(memberUrl, response)
}

func fetch(url string, response interface{}) error {
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
