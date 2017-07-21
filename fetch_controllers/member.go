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

type memberFetchResult struct {
	Members []models.Member `json:"results"`
}

type MemberFetchController struct {
	Senators        *[]models.Member
	Representatives *[]models.Member
}

func (mfc *MemberFetchController) InitialFetch(ft workers.FetchType, congress int) (apiUrls *[]string, error error) {
	var ifr initialFetchResult
	err := workers.PropublicaMembersFetch(congress, ft, &ifr)
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

func (mfc *MemberFetchController) FetchMember(url string) (*[]models.Member, error) {
	var mfr memberFetchResult
	err := workers.PropublicaMemberFetch(url, &mfr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return nil, err
	}

	return &(mfr.Members), nil
}

func (mfc *MemberFetchController) SaveMember(mem *models.Member, ct db.CollectionType) error {
	err := db.Insert(mem, ct)
	if err != nil {
		logger.Error("Failed to save member:", mem.FullName())
		return err
	}

	logger.Log("Successfully saved member:", mem.FullName())
	return nil
}

func (mfc *MemberFetchController) FetchAllMembers(ft workers.FetchType, congress int) error {
	urls, err := mfc.InitialFetch(ft, congress)
	if err != nil {
		logger.Error("Error: Failed to fetch member urls from propublica with error:", err)
		return err
	}

	members := make([]models.Member, 0, len(*urls))
	for _, url := range *urls {
		mems, err := mfc.FetchMember(url)
		if err != nil {
			logger.Warn("Failed to fetch from", ft, "at url:", url, "error:", err)
			continue
		}

		for _, m := range *mems {
			members = append(members, m)
			logger.Log("Fetched:", m.FullName(), "from", ft, "for congress:", congress)
		}
	}

	switch ft {
	case workers.HouseFetch:
		mfc.Representatives = &members
	case workers.SenateFetch:
		mfc.Senators = &members
	}

	return nil
}

func (mfc *MemberFetchController) SaveAllMembers(ft workers.FetchType) error {
	if ft == workers.HouseFetch && mfc.Representatives == nil {
		logger.Error("Can't save representatives. No representatives have been saved to the member fetch controller.")
		return errors.New("No Representatives To Save")
	}

	if ft == workers.SenateFetch && mfc.Senators == nil {
		logger.Error("Can't save senators. No senators have been saved to the member fetch controller.")
		return errors.New("No Senators To Save")
	}

	var ct db.CollectionType
	var members *[]models.Member
	switch ft {
	case workers.HouseFetch:
		members = mfc.Representatives
		ct = db.HouseCollection
	case workers.SenateFetch:
		members = mfc.Senators
		ct = db.SenateCollection
	}

	if len(*members) == 0 {
		logger.Warn(
			"Can't Save member for:",
			ft,
			"No member of that type have been saved to the member fetch controller.",
		)
		return nil
	}

	for _, mem := range *members {
		err := mfc.SaveMember(&mem, ct)
		if err != nil {
			logger.Warn("Failed to save member:", mem.FullName(), "to:", ft)
		}
	}

	return nil
}

func (mfc *MemberFetchController) FetchAndSaveMembers(ft workers.FetchType, congress int) error {
	err := mfc.FetchAllMembers(ft, congress)
	if err != nil {
		logger.Error("Failed to fetch and save members with fetch error:", err)
		return err
	}

	err = mfc.SaveAllMembers(ft)
	if err != nil {
		logger.Error("Failed to fetch and save members with save error:", err)
		return err
	}

	return nil
}

func (mfc *MemberFetchController) FetchAndSaveSenators(congress int) error {
	return mfc.FetchAndSaveMembers(workers.SenateFetch, congress)
}

func (mfc *MemberFetchController) FetchAndSaveRepresentatives(congress int) error {
	return mfc.FetchAndSaveMembers(workers.HouseFetch, congress)
}

func (mfc *MemberFetchController) FetchAndSaveSenatorsAndRepresentatives(congress int) error {
	err := mfc.FetchAndSaveSenators(congress)
	if err != nil {
		return err
	}

	return mfc.FetchAndSaveRepresentatives(congress)
}
