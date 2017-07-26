package fetchControllers

import (
	"errors"
	"github.com/kapitol-app/octopus/db"
	"github.com/kapitol-app/octopus/fetchers"
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/models"
)

type initialMemberFetchResult struct {
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

func (mfc *MemberFetchController) InitialFetch(ch fetchers.Chamber, congress int) (apiUrls *[]string, error error) {
	var ifr initialMemberFetchResult
	err := fetchers.PropublicaMembersFetch(congress, ch, &ifr)
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

	logger.Log("Retrieved:", len(urls), "urls for members of the", ch)
	return &urls, nil
}

func (mfc *MemberFetchController) FetchMember(url string) (*[]models.Member, error) {
	var mfr memberFetchResult
	err := fetchers.PropublicaMemberFetch(url, &mfr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return nil, err
	}

	return &(mfr.Members), nil
}

func (mfc *MemberFetchController) SaveMember(mem *models.Member, ct db.CollectionType) error {
	q := map[string]string{"memberId": mem.MemberId}
	err := db.Upsert(mem, &q, ct)
	if err != nil {
		logger.Error("Failed to save member:", mem.FullName())
		return err
	}

	logger.Log("Successfully saved member:", mem.FullName())
	return nil
}

func (mfc *MemberFetchController) FetchAllMembers(ch fetchers.Chamber, congress int) error {
	urls, err := mfc.InitialFetch(ch, congress)
	if err != nil {
		logger.Error("Error: Failed to fetch member urls from propublica with error:", err)
		return err
	}

	members := make([]models.Member, 0, len(*urls))
	for _, url := range *urls {
		mems, err := mfc.FetchMember(url)
		if err != nil {
			logger.Warn("Failed to fetch from", ch, "at url:", url, "error:", err)
			continue
		}

		for _, m := range *mems {
			members = append(members, m)
			logger.Log("Fetched:", m.FullName(), "from", ch, "for congress:", congress)
		}
	}

	switch ch {
	case fetchers.House:
		mfc.Representatives = &members
	case fetchers.Senate:
		mfc.Senators = &members
	}

	return nil
}

func (mfc *MemberFetchController) SaveAllMembers(ch fetchers.Chamber) error {
	if ch == fetchers.House && mfc.Representatives == nil {
		logger.Error("Can't save representatives. No representatives have been saved to the member fetch controller.")
		return errors.New("No Representatives To Save")
	}

	if ch == fetchers.Senate && mfc.Senators == nil {
		logger.Error("Can't save senators. No senators have been saved to the member fetch controller.")
		return errors.New("No Senators To Save")
	}

	var ct db.CollectionType
	var members *[]models.Member
	switch ch {
	case fetchers.House:
		members = mfc.Representatives
		ct = db.HouseCollection
	case fetchers.Senate:
		members = mfc.Senators
		ct = db.SenateCollection
	}

	if len(*members) == 0 {
		logger.Warn(
			"Can't Save member for:",
			ch,
			"No member of that type have been saved to the member fetch controller.",
		)
		return nil
	}

	for _, mem := range *members {
		err := mfc.SaveMember(&mem, ct)
		if err != nil {
			logger.Warn("Failed to save member:", mem.FullName(), "to:", ch)
		}
	}

	return nil
}

func (mfc *MemberFetchController) FetchAndSaveMembers(ch fetchers.Chamber, congress int) error {
	err := mfc.FetchAllMembers(ch, congress)
	if err != nil {
		logger.Error("Failed to fetch and save members with fetch error:", err)
		return err
	}

	err = mfc.SaveAllMembers(ch)
	if err != nil {
		logger.Error("Failed to fetch and save members with save error:", err)
		return err
	}

	return nil
}

func (mfc *MemberFetchController) FetchAndSaveSenators(congress int) error {
	return mfc.FetchAndSaveMembers(fetchers.Senate, congress)
}

func (mfc *MemberFetchController) FetchAndSaveRepresentatives(congress int) error {
	return mfc.FetchAndSaveMembers(fetchers.House, congress)
}

func (mfc *MemberFetchController) FetchAndSaveSenatorsAndRepresentatives(congress int) error {
	err := mfc.FetchAndSaveSenators(congress)
	if err != nil {
		return err
	}

	return mfc.FetchAndSaveRepresentatives(congress)
}
