package main

import (
	"github.com/kapitol-app/octopus/fetch_controllers"
	"github.com/kapitol-app/octopus/logger"
)

func main() {
	sfc := fetch_controllers.SenatorFetchController{}
	senators, err := sfc.FetchAllSenators()
	if err != nil {
		logger.Error("Failed to fetch all senators")
		return
	}

	logger.Log("Fetched:", len(*senators), "senators")
}
