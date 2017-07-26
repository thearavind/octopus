package fetchControllers

import (
	"github.com/kapitol-app/octopus/db"
	"github.com/kapitol-app/octopus/fetchers"
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/models"
)

type initialBillFetchResult struct {
	Results []struct {
		Bills []struct {
			BillUri string `json:"bill_uri"`
		} `json:"bills"`
	} `json:"results"`
}

type billFetchResult struct {
	Bills []models.Bill `json:"results"`
}

type voteFetchResult struct {
	Results struct {
		Votes struct {
			Vote models.Vote `json:"vote"`
		} `json:"votes"`
	} `json:"results"`
}

type BillFetchController struct {
	Bills *[]models.Bill
}

func (bfc *BillFetchController) FetchBills(congress int, ch fetchers.Chamber, bt fetchers.BillType) error {
	var ifr initialBillFetchResult
	err := fetchers.PropublicaBillFetch(congress, ch, bt, &ifr)
	if err != nil {
		logger.Log("Failed to to fetch all", ch, "bills with error:", err)
		return err
	}

	bills := make([]models.Bill, 0)
	for _, r := range ifr.Results {
		for _, m := range r.Bills {
			bill, err := bfc.FetchBillDetails(m.BillUri)
			if err != nil {
				logger.Log("Failed to fetch bill details", err)
				return err
			}

			logger.Log("Fetched:", bill.Bills[0].Title, "from", ch, "for congress:", congress)
			if len(bill.Bills[0].VotesUrl) > 0 {
				for _, b := range bill.Bills[0].VotesUrl {
					vote, err := bfc.FetchVoteDetails(b.APIURL)
					if err != nil {
						logger.Log("Failed to fetch vote details", err)
						return err
					}
					logger.Log("Fetched", vote.Chamber, "vote details for the bill", vote.Bill.Title)
					bill.Bills[0].Votes = append(bill.Bills[0].Votes, *vote)
				}
				bills = append(bills, bill.Bills[0])
			} else {
				logger.Log("No vote details available for bill", bill.Bills[0].Title)
				bills = append(bills, bill.Bills[0])
			}
		}
	}
	bfc.Bills = &bills
	return nil
}

func (bfc *BillFetchController) FetchBillDetails(billUrl string) (*billFetchResult, error) {
	br := billFetchResult{}
	err := fetchers.PropublicaMemberOrVoteFetch(billUrl, &br)
	if err != nil {
		logger.Log("Error: Failed to fetch bill details with error:", err)
		return nil, err
	}

	return &br, nil
}

func (bfc *BillFetchController) FetchVoteDetails(voteUrl string) (*models.Vote, error) {
	vr := voteFetchResult{}
	err := fetchers.PropublicaMemberOrVoteFetch(voteUrl, &vr)
	if err != nil {
		logger.Log("Error: Failed to fetch votes with error:", err)
		return nil, err
	}

	return &(vr.Results.Votes.Vote), nil
}

func (bfc *BillFetchController) SaveBill(bill *models.Bill) error {
	return db.Insert(bill, db.LegislativeCollection)
}

func (bfc *BillFetchController) fetchAndSaveAllBills(
	congress int,
	ch fetchers.Chamber,
	billType fetchers.BillType,
) (int, error) {
	count := 0
	err := bfc.FetchBills(congress, ch, billType)
	if err != nil {
		logger.Error("Error: Failed to fetch bills with error:", err)
		return count, err
	}

	for _, bill := range *bfc.Bills {
		err := bfc.SaveBill(&bill)
		if err == nil {
			count += 1
		} else {
			logger.Warn("Failed to save bill:", bill.Title, "to the database")
		}
	}

	return count, nil
}

func (bfc *BillFetchController) FetchAndSaveRecentSenateBills(congress int) (int, error) {
	return bfc.fetchAndSaveAllBills(congress, fetchers.Senate, fetchers.Passed)
}

func (bfc *BillFetchController) FetchAndSaveRecentHouseBills(congress int) (int, error) {
	return bfc.fetchAndSaveAllBills(congress, fetchers.House, fetchers.Passed)
}

func (bfc *BillFetchController) FetchAndSaveRecentBills(congress int) (int, error) {
	sc, err := bfc.FetchAndSaveRecentSenateBills(congress)
	if err != nil {
		return sc, err
	}

	hc, err := bfc.FetchAndSaveRecentHouseBills(congress)
	if err != nil {
		return hc, err
	}

	return sc + hc, nil
}
