package main

import (
	"github.com/kapitol-app/octopus/db"
	"github.com/kapitol-app/octopus/logger"
	"github.com/kapitol-app/octopus/models"
	"github.com/kapitol-app/octopus/workers"
)

func main() {
	senators := models.MemberListResponse{}
	err := workers.Fetch(115, workers.Senate, workers.Member, ".json", &senators)
	if err != nil {
		logger.Log("Failed to fetch senators data", "Error: ", err)
	} else {
		logger.Log("Fetched the senators data", senators, ",saving it to the DB")

		for i, senatorsLength := 0, senators.Results[0].NumResults; i < senatorsLength; i++ {
			err = db.Connection.SenatorCollection.Insert(&senators.Results[0].Members[i])
			if err != nil {
				logger.Log("Failed to insert the senator", senators.Results[0].Members[i], "into the DB")
				break
			}
		}
		logger.Log("Saved the senators data to the DB")
	}

	house := models.MemberListResponse{}
	err = workers.Fetch(115, workers.House, workers.Member, ".json", &house)
	if err != nil {
		logger.Log("Failed to fetch house members data", "Error: ", err)
	} else {
		logger.Log("Fetched the house members data", house, ",saving it to the DB")

		houseLength := house.Results[0].NumResults
		for i := 0; i < houseLength; i++ {
			err = db.Connection.MemberCollection.Insert(&house.Results[0].Members[i])
			if err != nil {
				logger.Log("Failed to insert the house member", house.Results[0].Members[i], "into the DB")
				break
			}
		}
		logger.Log("Saved the house members data to the DB")
	}
}
