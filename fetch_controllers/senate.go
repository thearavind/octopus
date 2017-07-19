package fetch_controllers

import (
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/models"
	"github.com/kapitol-app/octopus/workers"
)

type initialFetchResult struct {
	Results []struct {
		Members []struct {
			ApiUri string `json:"api_uri"`
		} `json:"members"`
	} `json:"results"`
}

type senatorFetchResult struct {
	Senators []models.Member `json:"results"`
}

type SenatorFetchController struct{}

func (sfc *SenatorFetchController) InitialFetch() (apiUrls *[]string, error error) {
	var ifr initialFetchResult
	err := workers.PropublicaMembersFetch(115, workers.Senate, &ifr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return nil, err
	}

	urls := make([]string, 0)
	for _, r := range ifr.Results {
		for _, m := range r.Members {
			urls = append(urls, string(m.ApiUri))
		}
	}

	return &urls, nil
}

func (sfc *SenatorFetchController) FetchSenator(url string) (*[]models.Member, error) {
	var sfr senatorFetchResult
	err := workers.PropublicaMemberFetch(url, &sfr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return nil, err
	}

	return &(sfr.Senators), nil
}
