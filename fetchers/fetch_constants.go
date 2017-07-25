package fetchers

//Chamber - Chamber of the congress
type Chamber string

const (
	House  Chamber = "house"
	Senate Chamber = "senate"
)

type BillType string

const (
	Introduced BillType = "introduced"
	Updated    BillType = "updated"
	Passed     BillType = "passed"
	Enacted    BillType = "enacted"
	Vetoed     BillType = "vetoed"
)
