package policy

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

const BASE_URL = "https://management.azure.com"

var checkForEmptyPath = regexp.MustCompile(`//`)

type Policy struct {
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	QueryValues map[string]string `json:"query_values"`
	PathValues  map[string]string `json:"path_values"`

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
	// add auth header
	req.Header.Add("Authorization", "Bearer "+token)

	return req, nil
}

// BuildRequestUrl builds the url for the request by using the Path template,
// and the PathValues map.
func (p Policy) BuildRequestUrl() (string, error) {
	// convert extra path values to PathValues map. This will allow
	// us to build the request url using a template.
	for k, v := range p.PathValues {
		p.PathValues[k] = v
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
