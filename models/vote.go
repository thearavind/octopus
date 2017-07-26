package models

import "reflect"


type Vote struct {
	Bill struct {
		APIURI       string `json:"api_uri" bson:"apiUri"`
		BillID       string `json:"bill_id" bson:"billId"`
		LatestAction string `json:"latest_action" bson:"latestAction"`
		Title        string `json:"title" bson:"title"`
	} `json:"bill" bson:"bill"`
	Date       string `json:"date" bson:"date"`
	Democratic struct {
		MajorityPosition string `json:"majority_position" bson:"majorityPosition"`
		No               int64  `json:"no" bson:"no"`
		NotVoting        int64  `json:"not_voting" bson:"notVoting"`
		Present          int64  `json:"present" bson:"present"`
		Yes              int64  `json:"yes" bson:"yes"`
	} `json:"democratic" bson:"democratic"`
	Description string `json:"description" bson:"description"`
	Independent struct {
		No        int64  `json:"no" bson:"no"`
		NotVoting string `json:"not_voting" bson:"notVoting"`
		Present   int64  `json:"present" bson:"present"`
		Yes       int64  `json:"yes" bson:"yes"`
	} `json:"independent" bson:"independent"`
	Positions []struct {
		DwNominate   float64 `json:"dw_nominate" bson:"dwNominate"`
		MemberID     string  `json:"member_id" bson:"memberId"`
		VotePosition string  `json:"vote_position" bson:"votePosition"`
	} `json:"positions" bson:"positions"`
	Question   string `json:"question" bson:"question"`
	Republican struct {
		MajorityPosition string `json:"majority_position" bson:"majorityPosition"`
		No               int64  `json:"no" bson:"no"`
		NotVoting        int64  `json:"not_voting" bson:"notVoting"`
		Present          int64  `json:"present" bson:"present"`
		Yes              int64  `json:"yes" bson:"yes"`
	} `json:"republican" bson:"republican"`
	Chamber        string `json:"chamber" bson:"chamber"`
	Result         string `json:"result" bson:"result"`
	RollCall       int64  `json:"roll_call" bson:"rollCall"`
	Session        int64  `json:"session" bson:"session"`
	TieBreaker     string `json:"tie_breaker" bson:"tieBreaker"`
	TieBreakerVote string `json:"tie_breaker_vote" bson:"tieBreakerVote"`
	Time           string `json:"time" bson:"time"`
	Total          struct {
		No        int64  `json:"no" bson:"no"`
		NotVoting string `json:"not_voting" bson:"notVoting"`
		Present   int64  `json:"present" bson:"present"`
		Yes       int64  `json:"yes" bson:"yes"`
	} `json:"total" bson:"total"`
	VoteType string `json:"vote_type" bson:"voteType"`
}

func (vote *Vote) Equals(v *Vote) bool {
	return reflect.DeepEqual(vote, v)
}
