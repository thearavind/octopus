package models

import "time"

type Committee struct {
	Name        string    `json:"name" bson:"name"`
	Code        string    `json:"code" bson:"code"`
	ApiUri      string    `json:"api_uri" bson:"apiUri"`
	RankInParty string    `json:"rank_in_party" bson:"rankInParty"`
	BeginDate   time.Time `json:"begin_date" bson:"beginDate"`
	EndDate     time.Time `json:"end_date" bson:"endDate"`
}
