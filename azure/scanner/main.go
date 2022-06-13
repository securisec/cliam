package scanner

import (
	"io/ioutil"
	"net/http"
)

const USER_AGENT = "AZURECLI/2.37.0 (HOMEBREW) azsdk-python-azure-mgmt-web/6.1.0 Python/3.10.4"

// MakeRequest makes a request to the policy's rest api.
func MakeRequest(req *http.Request) (int, []byte, error) {
	req.Header.Add("User-Agent", USER_AGENT)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	return res.StatusCode, body, nil
}
