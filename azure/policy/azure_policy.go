package policy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"text/template"
)

const BASE_URL = "https://management.azure.com"

var checkForEmptyPath = regexp.MustCompile(`//`)

type Policy struct {
	Path          string                 `json:"path"`
	Method        string                 `json:"method"`
	QueryValues   map[string]string      `json:"query_values"`
	HeadersValues map[string]string      `json:"headers_values"`
	JsonVaues     map[string]interface{} `json:"json_values"`
	FormValues    map[string]string      `json:"form_values"`
	PathValues    map[string]string      `json:"path_values"`
	IsExtra       bool                   `json:"is_extra"`
	ExtraLocation string                 `json:"extra_location"`
	ExtraValue    map[string]interface{} `json:"extra_value"`

	// extra
	OperationID string `json:"operation_id"`
	Resource    string `json:"resource"`
}

type CommonPathProperties struct {
	SubscriptionID    string
	ResourceGroupName string
}

func (p Policy) BuildRequest(ctx context.Context, token string, common CommonPathProperties) (*http.Request, error) {
	body := strings.NewReader("")
	method := "GET"
	if p.Method != "" {
		method = p.Method
	}

	if p.PathValues == nil {
		p.PathValues = make(map[string]string)
	}
	if common.SubscriptionID != "" {
		p.PathValues["subscriptionId"] = common.SubscriptionID
	}
	if common.ResourceGroupName != "" {
		p.PathValues["resourceGroupName"] = common.ResourceGroupName
	}

	// build url
	u, err := p.BuildRequestUrl()
	if err != nil {
		return nil, err
	}

	// handle body
	// if form data, set body
	if method != "GET" && len(p.FormValues) > 0 {
		form := url.Values{}
		for k, v := range p.FormValues {
			form.Add(k, v)
		}
		body = strings.NewReader(form.Encode())
	}

	if method != "GET" && len(p.FormValues) == 0 {
		o, err := json.Marshal(p.JsonVaues)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(o))
	}
	// build request
	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, err
	}

	// add query values
	q := req.URL.Query()
	for k, v := range p.QueryValues {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// add headers
	for k, v := range p.HeadersValues {
		req.Header.Add(k, v)
	}
	// add auth header
	req.Header.Add("Authorization", "Bearer "+token)

	// build extra
	if p.IsExtra {
		req, err = p.BuildExtra(req)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}

// BuildRequestUrl builds the url for the request by using the Path template,
// and the PathValues map.
func (p Policy) BuildRequestUrl() (string, error) {
	// convert extra path values to PathValues map. This will allow
	// us to build the request url using a template.
	if p.IsExtra && p.ExtraLocation == "path" {
		for k, v := range p.ExtraValue {
			v, ok := v.(string)
			if !ok {
				continue
			}
			p.PathValues[k] = v
		}
	}
	var b bytes.Buffer
	t := template.Must(template.New("url").Parse(p.Path))
	t = t.Option("missingkey=error")
	if err := t.Execute(&b, p.PathValues); err != nil {
		return "", err
	}

	// check for empty values
	if checkForEmptyPath.MatchString(b.String()) {
		return "", fmt.Errorf("empty value detected")
	}

	return BASE_URL + b.String(), nil
}

func (p Policy) BuildExtra(req *http.Request) (*http.Request, error) {
	// TODO 🚧
	switch p.ExtraLocation {
	// case condition:

	default:
		return req, nil
	}
}
