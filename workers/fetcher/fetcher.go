package fetcher

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/greenac/octopus/logger"
	"errors"
)

type Fetcher struct {
	Url string
}

func (f *Fetcher)Fetch() (contents []byte, error error) {
	resp, err := http.Get(f.Url)
	if err != nil {
		fmt.Println("Error: fetching contents from:", f.Url, "error:", err)
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: reading contents from:", f.Url, "error:", err)
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return content, nil
}

func (f *Fetcher)PostJson(body *[]byte) (responseData []byte, err error) {
	var data []byte
	resp, err := http.Post(f.Url, "application/json", body)
	if err != nil {
		logger.Log("Error: could not post json to:", f.Url)
		return data, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching json from:", f.Url, "Got status:", resp.StatusCode)
		return data, errors.New("StatusCodeNotOk")
	}

	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}
