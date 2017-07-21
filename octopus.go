package main

import (
	"github.com/kapitol-app/octopus/fetch_controllers"
)

func main() {
	mfc := fetch_controllers.MemberFetchController{}
	mfc.FetchAndSaveSenatorsAndRepresentatives(115)
}
