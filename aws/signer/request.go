package signer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/securisec/cliam/aws/policy"
	"github.com/securisec/cliam/logger"
)

func MakeRequest(
	ctx context.Context,
	region, service string,
	p *policy.Service,
	creds *credentials.Credentials,
	requestURL string,
) (*http.Request, *http.Response, []byte, error) {
	// default options
	method := "GET"
	body := strings.NewReader("")

	if p.Method != "" {
		method = p.Method
	}

	// if form data, set body
	if method != "GET" && len(p.FormData) > 0 {
		form := url.Values{}
		for k, v := range p.FormData {
			form.Add(k, v)
		}
		body = strings.NewReader(form.Encode())
	}

	if method != "GET" && len(p.FormData) == 0 {
		o, err := json.Marshal(p.JsonData)
		if err != nil {
			return nil, nil, nil, err
		}
		body = strings.NewReader(string(o))
	}
	// create request
	req, err := http.NewRequestWithContext(ctx, method, requestURL, body)
	if err != nil {
		return nil, nil, nil, err
	}

	// add headers
	if p.Headers != nil {
		for k, v := range p.Headers {
			req.Header.Add(k, v)
		}
	}

	// we only want accent json
	req.Header.Add("Accept", "application/json")

	// create signer to generate auth signature header
	signer := v4.NewSigner(creds)
	// safety net for certain services
	service = serviceSafetyNet(service)
	// safety net for certain regions
	region = policy.AwsRegionSafetyNet(service, region)
	logger.LogDebugVerbose("url", requestURL)
	// sign request
	if _, err := signer.Sign(req, body, service, region, time.Now()); err != nil {
		return nil, nil, nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, nil, err
	}
	res.Body.Close()

	if logger.DEBUG && logger.VERBOSE {
		logger.LoggerStdErr.Debug().Msg(string(b))
	}

	return req, res, b, nil
}

func serviceSafetyNet(service string) string {
	switch service {
	case "streams.dynamodb":
		return "dynamodb"
	default:
		return service
	}
}
