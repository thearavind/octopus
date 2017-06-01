package workers

import(
	"net/http"
	"io/ioutil"
	"github.com/kapitol-app/octopus/logger"
	//"errors"
)

type Fetcher struct {
	Url string
}

func (f *Fetcher)Fetch() (contents []byte, error error) {
	resp, err := http.Get(f.Url)
	if err != nil {
		logger.Log("Error: fetching contents from:", f.Url, "error:", err)
		return nil, err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log("Error: reading contents from:", f.Url, "error:", err)
		return nil, err
	}

	return content, nil
}

//func (f *Fetcher)PostJson(body *[]byte) (responseData []byte, err error) {
//	var data []byte
//	resp, err := http.Post(f.Url, "application/json", body)
//	if err != nil {
//		logger.Log("Error: could not post json to:", f.Url)
//		return data, err
//	}
//
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		logger.Log("Error fetching json from:", f.Url, "Got status:", resp.StatusCode)
//		return data, errors.New("StatusCodeNotOk")
//	}
//
//	data, err = ioutil.ReadAll(resp.Body)
//	return data, err
//}
