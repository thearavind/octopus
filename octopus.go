package main

import (
	"github.com/kapitol-app/octopus/fetchControllers"
	"github.com/kapitol-app/octopus/logger"
)

func main() {
	mfc := fetchControllers.MemberFetchController{}
	err := mfc.FetchAndSaveSenatorsAndRepresentatives(115)
	if err == nil {
		logger.Log("Successfully fetched and saved senators and representatives")
	} else {
		logger.Error("Failed to fetch and save senators and representatives")
	}

	bfc := fetchControllers.BillFetchController{}
	successCount, err := bfc.FetchAndSaveRecentBills(115)
	if err == nil {
		logger.Log("Fetched and saved:", successCount, "recent bills from the house and senate")
	} else {
		logger.Error("Failed to fetch and save all recent senate bills")
	}
}
