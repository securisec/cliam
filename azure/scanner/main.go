package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const USER_AGENT = "AZURECLI/2.37.0 (HOMEBREW) azsdk-python-azure-mgmt-web/6.1.0 Python/3.10.4"

// MakeRequest makes a request to the policy's rest api.
func MakeRequest(req *http.Request) (*http.Response, map[string]interface{}, error) {
	req.Header.Add("User-Agent", USER_AGENT)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	// unmarshal body into a map
	var b map[string]interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, nil, fmt.Errorf(string(body))
	}
	// check if error is in the body
	if _, ok := b["error"]; ok {
		return nil, nil, fmt.Errorf("%s", b["error"])
	}
	return res, b, nil
}
