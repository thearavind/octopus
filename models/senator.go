package models

//Senator - Senators model
type Senator struct {
	APIURI          string `json:"api_uri" bson:"apiUri"`
	FacebookAccount string `json:"facebook_account" bson:"facebookAccount"`
	FirstName       string `json:"first_name" bson:"firstName"`
	GovtrackID      string `json:"govtrack_id" bson:"govTrackId"`
	ID              string `json:"id" bson:"ppId"`
	InOffice        string `json:"in_office" bson:"inOffice"`
	LastName        string `json:"last_name" bson:"lastName"`
	LeadershipRole  string `json:"leadership_role" bson:"leadershipRole"`
	LisID           string `json:"lis_id" bson:"lisId"`
	MissedVotes     int    `json:"missed_votes" bson:"missedVotes"`
	NextElection    string `json:"next_election" bson:"nextElection"`
	Office          string `json:"office" bson:"office"`
	Party           string `json:"party" bson:"party"`
	Phone           string `json:"phone" bson:"phone"`
	SenateClass     string `json:"senate_class" bson:"senateClass"`
	State           string `json:"state" bson:"state"`
	StateRank       string `json:"state_rank" bson:"stateRank"`
	TotalVotes      int    `json:"total_votes" bson:"totalVotes"`
	TwitterAccount  string `json:"twitter_account" bson:"twitterAccount"`
	URL             string `json:"url" bson:"url"`
}
