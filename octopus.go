package main

import (
	"fmt"
	"github.com/kapitol-app/octopus/fetch_controllers"
	"github.com/kapitol-app/octopus/logger"
)

func main() {
	sfc := fetch_controllers.SenatorFetchController{}
	urls, err := sfc.InitialFetch()
	if err != nil {
		logger.Log("Error: Failed to fetch urls from propublica with error:", err)
		return
	}

	url := (*urls)[0]
	sens, err := sfc.FetchSenator(url)
	if err != nil {
		logger.Log("Failed to fetch senator:", url, "error:", err)
		return
	}

	for _, s := range *sens {
		fmt.Println("senator:", s.FullName())
		fmt.Println("roles:")
		for _, r := range *s.Roles {
			fmt.Println(r.Title)
			fmt.Println("committees:")
			for _, c := range *r.Committees {
				fmt.Println(c.Name)
			}
		}
	}
}
