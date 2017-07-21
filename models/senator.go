package models

// TODO: When we finish the House model, both the Senator Model and
// HouseMember Model should Inherit from the Member model. At that
// time we should move any like functions and field into the base model.
// Leaving it this way for now because I'm lazy.

type Senator struct {
	MemberId        string  `json:"member_id" bson:"memberId"`
	FirstName       string  `json:"first_name" bson:"firstName"`
	MiddleName      string  `json:"middle_name" bson:"middleName"`
	LastName        string  `json:"last_name" bson:"lastName"`
	DateOfBirth     string  `json:"date_of_birth" bson:"dateOfBirth"`
	Gender          string  `json:"gender" bson:"gender"`
	Url             string  `json:"url" bson:"url"`
	TimesTopicUrl   string  `json:"times_topics_url" bson:"timesTopicsUrl"`
	TimesTag        string  `json:"times_tag" bson:"timesTag"`
	GovtrackId      string  `json:"govtrack_id" bson:"govtrackId"`
	CspanId         string  `json:"cspan_id" bson:"cspanId"`
	VotesmartId     int     `json:"votesmart_id,string" bson:"votesmartId"`
	IcpsrId         string  `json:"icpsr_id" bson:"icpsrId"`
	TwitterAccount  string  `json:"twitter_account" bson:"twitterAccount"`
	FacebookAccount string  `json:"facebook_account" bson:"facebookAccount"`
	YoutubeAccount  string  `json:"youtube_account" bson:"youtubeAccount"`
	CrpId           string  `json:"crp_id" bson:"crpId"`
	GoogleEntityId  string  `json:"google_entity_id" bson:"googlEntityId"`
	RssUrl          string  `json:"rss_url" bson:"rssUrl"`
	InOffice        bool    `json:"in_office" bson:"inOffice"`
	CurrentParty    string  `json:"current_party" bson:"currentParty"`
	MostRecentVote  string  `json:"most_recent_vote" bson:"mostRecentVote"`
	Roles           *[]Role `json:"roles" bson:"roles"`
}

func (m *Senator) FullName() string {
	return m.FirstName + " " + m.LastName
}
