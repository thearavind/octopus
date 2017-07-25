package fetchControllers

import (
	"github.com/kapitol-app/octopus/db"
	"github.com/kapitol-app/octopus/fetchers"
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/models"
)

type initialBillFetchResult struct {
	Results []struct {
		Bills *[]models.Bill `json:"bills"`
	} `json:"results"`
}

type BillFetchController struct {
	Bills *[]models.Bill
}

func (bfc *BillFetchController) FetchBills(congress int, ch fetchers.Chamber, bt fetchers.BillType) error {
	var ifr initialBillFetchResult
	err := fetchers.PropublicaBillFetch(congress, ch, bt, &ifr)
	if err != nil {
		logger.Log("Failed to to fetch all senate members with error:", err)
		return err
	}

	bills := make([]models.Bill, 0)
	for _, r := range ifr.Results {
		for _, bill := range *r.Bills {
			bills = append(bills, bill)
		}
	}

	bfc.Bills = &bills
	return nil
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
