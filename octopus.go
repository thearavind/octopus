package main

import (
	"github.com/kapitol-app/octopus/fetch_controllers"
	"github.com/kapitol-app/octopus/logger"
)

func main() {
	mfc := fetch_controllers.MemberFetchController{}
	err := mfc.FetchAndSaveSenatorsAndRepresentatives(115)
	if err == nil {
		logger.Log("Successfully fetched and saved senators and representatives")
	} else {
		logger.Error("Failed to fetch and save senators and representatives")
	}
}
