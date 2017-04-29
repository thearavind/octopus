package models

type Model interface {
	toMap() map[string]interface{}
}

type Member struct {
	PpId string `json:"id"`
	ApiUri string `json:"api_uri"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Party string `json:"party"`
	LeadershipRole string `json:"leadership_role"`
	TwitterAccount string `json:"leadership_role"`
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
	m["PpId"] = mem.PpId
	m["ApiUri"] = mem.ApiUri
	m["FirstName"] = mem.FirstName
	m["LastName"] = mem.LastName
	m["Party"] = mem.Party
	m["LeadershipRole"] = mem.LeadershipRole
	m["TwitterAccount"] = mem.TwitterAccount
	m["Url"] = mem.Url
	m["GovTrackId"] = mem.GovTrackId
	m["InOffice"] = mem.InOffice
	m["NextElection"] = mem.NextElection
	m["TotalVotes"] = mem.TotalVotes
	m["MissedVotes"] = mem.MissedVotes
	m["Office"] = mem.Office
	m["Phone"] = mem.Phone
	m["State"] = mem.State

	return m
}

type Senator struct {
	Member
	SenateClass string `json:"senate_class"`
	MissedVotesPct string `json:"missed_votes_pct"`
	VotesWithPartyPct string `json:"votes_with_party_pct"`
}

func (s *Senator)toMap() map[string]interface{} {
	m := s.Member.toMap()
	m["SenateClass"] = s.SenateClass
	m["MissedVotesPct"] = s.MissedVotesPct
	m["VotesWithPartyPct"] = s.VotesWithPartyPct

	return m
}
