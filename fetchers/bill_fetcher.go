package fetchers

import (
	"fmt"
	"github.com/kapitol-app/octopus/endpoints"
)

func PropublicaBillFetch(congress int, ch Chamber, bt BillType, response interface{}) error {
	url := fmt.Sprintf(
		"%s%d/%s/bills/%s.json",
		endpoints.GetEndpoint(endpoints.Propublica),
		congress,
		ch,
		bt,
	)

	return Fetch(url, response)
}
