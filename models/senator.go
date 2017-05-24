package models

type Senator struct {
	Member
	SenateClass string `json:"senate_class"`
	MissedVotesPct float64 `json:"missed_votes_pct"`
	VotesWithPartyPct float64 `json:"votes_with_party_pct"`
}

func (s *Senator)toMap() map[string]interface{} {
	m := s.Member.toMap()
	m["SenateClass"] = s.SenateClass
	m["MissedVotesPct"] = s.MissedVotesPct
	m["VotesWithPartyPct"] = s.VotesWithPartyPct

	return m
}
