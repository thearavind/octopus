package models

import "reflect"

type Role struct {
	Congress          int     `json:"congress,string" bson:"congress"`
	Chamber           string  `json:"chamber" bson:"chamber"`
	Title             string  `json:"title" bson:"title"`
	State             string  `json:"state" bson:"state"`
	Party             string  `json:"party" bson:"party"`
	LeadershipRole    string  `json:"leadership_role" bson:"leadershipRole"`
	FecCandidateId    string  `json:"fec_candidate_id" bson:"fecCandidateId"`
	Seniority         int     `json:"seniority,string" bson:"seniority"`
	SenateClass       int     `json:"senate_class,string" bson:"senateClass"`
	District          int     `json:"district,string" bson:"district,omitempty"`
	AtLarge           bool    `json:"at_large" bson:"atLarge"`
	StateRank         string  `json:"state_rank" bson:"stateRank"`
	LisId             string  `json:"lis_id" bson:"lisId"`
	OcdId             string  `json:"ocd_id" bson:"ocdId"`
	StartDate         string  `json:"start_date" bson:"startDate"`
	EndDate           string  `json:"end_date" bson:"endDate"`
	Office            string  `json:"office" bson:"office"`
	Phone             string  `json:"phone" bson:"phone"`
	Fax               string  `json:"fax" bson:"fax"`
	ContactForm       string  `json:"contact_form" bson:"contactForm"`
	BillsSponsored    int     `json:"bills_sponsored" bson:"billsSponsored"`
	BillsCosponsored  int     `json:"bills_cosponsored" bson:"billsCosponsored"`
	MissedVotesPct    float64 `json:"missed_votes_pct" bson:"missedVotesPct"`
	VotesWithPartyPct float64 `json:"votes_with_party_pct" bson:"votesWithPartyPct"`
	Committees        *[]Committee
}

func (role *Role) Equals(r *Role) bool {
	return reflect.DeepEqual(role, r)
}
