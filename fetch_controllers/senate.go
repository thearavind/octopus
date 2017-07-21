package fetch_controllers

import (
	"errors"
	"github.com/kapitol-app/octopus/db"
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
	Senators []models.Senator `json:"results"`
}

type SenatorFetchController struct {
	Senators *[]models.Senator
}

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

func (sfc *SenatorFetchController) FetchSenator(url string) (*[]models.Senator, error) {
	var sfr senatorFetchResult
	err := workers.PropublicaMemberFetch(url, &sfr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return nil, err
	}

	return &(sfr.Senators), nil
}

func (sfc *SenatorFetchController) FetchAllSenators() error {
	urls, err := sfc.InitialFetch()
	if err != nil {
		logger.Error("Error: Failed to fetch urls from propublica with error:", err)
		return err
	}

	senators := make([]models.Senator, 0, len(*urls))
	for _, url := range *urls {
		sens, err := sfc.FetchSenator(url)
		if err != nil {
			logger.Warn("Failed to fetch senator:", url, "error:", err)
			continue
		}

		for _, s := range *sens {
			senators = append(senators, s)
			logger.Log("Fetched:", s.FullName())
		}
	}

	sfc.Senators = &senators
	return nil
}

func (sfc *SenatorFetchController) SaveSenator(sen *models.Senator) error {
	err := db.Insert(sen, db.SenateCollection)
	if err != nil {
		logger.Error("Failed to save senator:", sen.FullName())
		return err
	}

	logger.Log("Successfully saved senator:", sen.FullName())
	return nil
}

func (sfc *SenatorFetchController) SaveAllSenators() error {
	if sfc.Senators == nil {
		logger.Error("Can't Save Senators. No senators have been saved to the senate fetch controller.")
		return errors.New("No Senators To Save")
	}

	if len(*sfc.Senators) == 0 {
		logger.Warn("Can't Save Senators. No senators have been saved to the senate fetch controller.")
		return nil
	}

	for _, sen := range *sfc.Senators {
		err := sfc.SaveSenator(&sen)
		if err != nil {
			logger.Warn("Failed to save senator:", sen.FullName())
		}
	}

	return nil
}

func (sfc *SenatorFetchController) FetchAndSaveSenators() error {
	err := sfc.FetchAllSenators()
	if err != nil {
		logger.Error("Failed to fetch and save senators with fetch error:", err)
		return err
	}

	err = sfc.SaveAllSenators()
	if err != nil {
		logger.Error("Failed to fetch and save senators with save error:", err)
		return err
	}

	return nil
}
