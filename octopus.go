package main

import (
	"github.com/kapitol-app/octopus/fetch_controllers"
)

func main() {
	sfc := fetch_controllers.SenatorFetchController{}
	sfc.FetchAndSaveSenators()
}
