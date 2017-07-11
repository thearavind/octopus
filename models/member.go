package models

//Member - members model
type Member struct {
	APIURI          string `json:"api_uri" bson:"apiUri"`
	District        string `json:"district" bson:"district"`
	FacebookAccount string `json:"facebook_account" bson:"facebookAccount"`
	FirstName       string `json:"first_name" bson:"firstName"`
	Geoid           string `json:"geoid" bson:"geoId"`
	GovtrackID      string `json:"govtrack_id" bson:"govTrackId"`
	ID              string `json:"id" bson:"ppId"`
	InOffice        string `json:"in_office" bson:"inOffice"`
	LastName        string `json:"last_name" bson:"lastName"`
	LeadershipRole  string `json:"leadership_role" bson:"leadershipRole"`
	MissedVotes     int    `json:"missed_votes" bson:"missedVotes"`
	NextElection    string `json:"next_election" bson:"nextElection"`
	Office          string `json:"office" bson:"office"`
	Party           string `json:"party" bson:"party"`
	Phone           string `json:"phone" bson:"phone"`
	State           string `json:"state" bson:"state"`
	TotalVotes      int    `json:"total_votes" bson:"totalVotes"`
	TwitterAccount  string `json:"twitter_account" bson:"twitterAccount"`
	URL             string `json:"url" bson:"url"`
}
