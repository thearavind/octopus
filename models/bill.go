package models

type Bill struct {
	BillId                string    `json:"bill_id" bson:"billId"`
	BillType              string    `json:"bill_type" bson:"billType"`
	Number                string    `json:"number" bson:"number"`
	BillUri               string    `json:"bill_uri" bson:"billUri"`
	Title                 string    `json:"title" bson:"title"`
	SponsorId             string    `json:"sponsor_id" bson:"sponsorId"`
	SponsorName           string    `json:"sponsor_name" bson:"sponsorName"`
	SponsorState          string    `json:"sponsor_state" bson:"sponsorState"`
	SponsorParty          string    `json:"sponsor_party" bson:"sponsorParty"`
	SponsorUri            string    `json:"sponsor_uri" bson:"sponsorUri"`
	GpoPdfUri             string    `json:"gpo_pdf_uri" bson:"gpoPdfUri"`
	CongressDotGovUrl     string    `json:"congressdotgov_url" bson:"congressDotGovUrl"`
	GovTrackUrl           string    `json:"govtrack_url" bson:"govTrackUrl"`
	IntroducedDate        string    `json:"introduced_date" bson:"introducedDate"`
	Active                bool      `json:"next_election" bson:"nextElection"`
	HousePassage          string    `json:"house_passage" bson:"housePassage"`
	SenatePassage         string    `json:"senate_passage" bson:"senatePassage"`
	Enacted               string    `json:"enacted" bson:"enacted"`
	Vetoed                string    `json:"vetoed" bson:"vetoed"`
	Committees            string    `json:"committees" bson:"committees"`
	PrimarySubject        string    `json:"primary_subject" bson:"primarySubject"`
	Summary               string    `json:"summary" bson:"summary"`
	SummaryShort          string    `json:"summary_short" bson:"summaryShort"`
	LatestMajorActionDate string    `json:"latest_major_action_date" bson:"latestMajorActionDate"`
	LatestMajorAction     string    `json:"latest_major_action" bson:"latestMajorAction"`
	CommitteeCodes        *[]string `json:"committee_codes" bson:"committeeCodes"`
}
