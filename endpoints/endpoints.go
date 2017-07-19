package endpoints

import "github.com/kapitol-app/octopus/config"

type Endpoint string

const (
	Propublica Endpoint = "propublica"
)

func GetEndpoint(ep Endpoint) string {
	var endPoint string
	switch ep {
	case Propublica:
		endPoint = config.C.ApiUrls.Propublica
	default:
		panic("Error: provided endpoint is not a known end point")
	}

	return endPoint
}
