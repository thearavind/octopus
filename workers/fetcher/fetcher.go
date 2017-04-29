package fetcher

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/greenac/chpt1/models"
	"errors"
)

type Fetcher struct {
	Url string
}

func (f *Fetcher) Fetch() (contents []byte, error error) {
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

func (f *Fetcher) GetUser() (user.User, error) {
	var u user.User
	resp, err := http.Get(f.Url)
	if err != nil {
		fmt.Println("Error fetching json from:", f.Url, "error:", err)
		return u, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching json from:", f.Url, "Got status:", resp.StatusCode)
		resp.Body.Close()
		return u, errors.New("StatusCodeNotOk")
	}

	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		fmt.Println("Error decoding json from:", f.Url, "error:", err)
		resp.Body.Close()
		return u, err
	}

	resp.Body.Close()
	return u, nil
}
