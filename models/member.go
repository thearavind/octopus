package models

type Member struct {
	PpId string `json:"id"`
	ApiUri string `json:"api_uri"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Party string `json:"party"`
	LeadershipRole string `json:"leadership_role"`
	TwitterAccount string `json:"twitter_account"`
	FacebookAccount string `json:"facebook_account"`
	Url string `json:"url"`
	GovTrackId string `json:"govtrack_id"`
	InOffice bool `json:"in_office"`
	NextElection int `json:"next_election"`
	TotalVotes int `json:"total_votes"`
	MissedVotes int `json:"missed_votes"`
	Office string `json:"office"`
	Phone string `json:"phone"`
	State string `json:"state"`
}

func (mem *Member)toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["ppId"] = mem.PpId
	m["apiUri"] = mem.ApiUri
	m["firstName"] = mem.FirstName
	m["lastName"] = mem.LastName
	m["party"] = mem.Party
	m["leadershipRole"] = mem.LeadershipRole
	m["twitterAccount"] = mem.TwitterAccount
	m["facebookAccount"] = mem.FacebookAccount
	m["url"] = mem.Url
	m["govTrackId"] = mem.GovTrackId
	m["inOffice"] = mem.InOffice
	m["nextElection"] = mem.NextElection
	m["totalVotes"] = mem.TotalVotes
	m["missedVotes"] = mem.MissedVotes
	m["office"] = mem.Office
	m["phone"] = mem.Phone
	m["state"] = mem.State

	return m
}
