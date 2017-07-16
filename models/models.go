package models

//MemberListResponse - Memeber response format
type MemberListResponse struct {
	Copyright string `json:"copyright" bson:"-"`
	Results   []struct {
		Chamber  string `json:"chamber" bson:"-"`
		Congress string `json:"congress" bson:"-"`
		Members  []struct {
			APIURI            string  `json:"api_uri" bson:"apiUri"`
			CrpID             string  `json:"crp_id" bson:"-"`
			CspanID           string  `json:"cspan_id" bson:"-"`
			District          string  `json:"district" bson:"district,omitempty"`
			Domain            string  `json:"domain" bson:"-"`
			DwNominate        float64 `json:"dw_nominate" bson:"-"`
			FacebookAccount   string  `json:"facebook_account" bson:"facebookAccount"`
			FirstName         string  `json:"first_name" bson:"firstName"`
			Geoid             string  `json:"geoid" bson:"geoId,omitempty"`
			GoogleEntityID    string  `json:"google_entity_id" bson:"-"`
			GovtrackID        string  `json:"govtrack_id" bson:"govTrackId"`
			IcpsrID           string  `json:"icpsr_id" bson:"-"`
			ID                string  `json:"id" bson:"ppId"`
			IdealPoint        string  `json:"ideal_point" bson:"-"`
			InOffice          string  `json:"in_office,string" bson:"inOffice"`
			LastName          string  `json:"last_name" bson:"lastName"`
			LeadershipRole    string  `json:"leadership_role" bson:"leadershipRole"`
			LisID             string  `json:"lis_id" bson:"lisId,omitempty"`
			MiddleName        string  `json:"middle_name" bson:"-"`
			MissedVotes       int     `json:"missed_votes" bson:"missedVotes"`
			MissedVotesPct    float32 `json:"missed_votes_pct" bson:"-"`
			NextElection      string  `json:"next_election" bson:"nextElection"`
			OcdID             string  `json:"ocd_id" bson:"-"`
			Office            string  `json:"office" bson:"office"`
			Party             string  `json:"party" bson:"party"`
			Phone             string  `json:"phone" bson:"phone"`
			RssURL            string  `json:"rss_url" bson:"-"`
			SenateClass       string  `json:"senate_class" bson:"senateClass,omitempty"`
			Seniority         string  `json:"seniority" bson:"-"`
			State             string  `json:"state" bson:"state"`
			StateRank         string  `json:"state_rank" bson:"stateRank,omitempty"`
			TotalPresent      int     `json:"total_present" bson:"-"`
			TotalVotes        int     `json:"total_votes" bson:"totalVotes"`
			TwitterAccount    string  `json:"twitter_account" bson:"twitterAccount"`
			URL               string  `json:"url" bson:"url"`
			VotesWithPartyPct float64 `json:"votes_with_party_pct" bson:"-"`
			VotesmartID       string  `json:"votesmart_id" bson:"-"`
		} `json:"members"`
		NumResults int `json:"num_results" bson:"-"`
		Offset     int `json:"offset" bson:"-"`
	} `json:"results"`
	Status string `json:"status" bson:"-"`
}
