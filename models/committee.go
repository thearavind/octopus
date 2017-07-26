package models

import "reflect"

type Committee struct {
	Name        string `json:"name" bson:"name"`
	Code        string `json:"code" bson:"code"`
	ApiUri      string `json:"api_uri" bson:"apiUri"`
	RankInParty int    `json:"rank_in_party,omitempty" bson:"rankInParty"`
	BeginDate   string `json:"begin_date" bson:"beginDate"`
	EndDate     string `json:"end_date" bson:"endDate"`
}

func (comm *Committee) Equals(c *Committee) bool {
	return reflect.DeepEqual(comm, c)
}
