package fetchers

import (
	"fmt"
	"github.com/kapitol-app/octopus/endpoints"
)

//Fetch - Fetches the data from the propublica api
func PropublicaMembersFetch(congress int, ch Chamber, response interface{}) error {
	url := fmt.Sprintf(
		"%s%d/%s/%s",
		endpoints.GetEndpoint(endpoints.Propublica),
		congress,
		ch,
		"members.json",
	)

	return Fetch(url, response)
}

func PropublicaMemberFetch(memberUrl string, response interface{}) error {
	return Fetch(memberUrl, response)
}
