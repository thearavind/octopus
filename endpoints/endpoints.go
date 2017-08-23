package endpoints

import "github.com/kapitol-app/octopus/config"

type Endpoint string

const (
	Propublica         Endpoint = "propublica"
	TwitterAccessToken Endpoint = "twitterToken"
	TwitterSearch      Endpoint = "twitterSearch"
)

func GetEndpoint(ep Endpoint) string {
	var endPoint string
	switch ep {
	case Propublica:
		endPoint = config.C.ApiUrls.Propublica
	case TwitterAccessToken:
		endPoint = config.C.ApiUrls.TwitterAccessToken
	case TwitterSearch:
		endPoint = config.C.ApiUrls.TwitterSearch
	default:
		panic("Error: provided endpoint is not a known end point")
	}

	return endPoint
}
