package gonjalla

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type request struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

const endpoint string = "https://njal.la/api/1/"

// HTTPClient interface. Useful for mocked unit tests later on.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client used in all requests by Request. Can be overwritten for tests.
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Request common function for all of Njalla's API.
// Njalla's API uses JSON-RPC, and contains just one endpoint.
// The endpoint is POST only, and takes in a JSON in the body, with two
// arguments, check the `request` struct for more info.
// The `params` argument is variable. Some methods require no parameters,
// (like `list-domains`), while other methods require parameters (like
// `get-domain` which requires `domain: string`).
func Request(
	token string, method string, params map[string]interface{},
) ([]byte, error) {
	token = fmt.Sprintf("Njalla %s", token)

	body, err := json.Marshal(
		request{Method: method, Params: params},
	)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	result, ok := data["result"]
	if !ok {
		return nil, fmt.Errorf("Missing result %s", data)
	}

	unwrapped, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return unwrapped, nil
}
