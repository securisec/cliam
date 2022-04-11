package signer

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/securisec/enumerate/logger"
	"github.com/securisec/enumerate/policy"
)

func MakeRequest(
	ctx context.Context,
	region, service string,
	p *policy.Service,
	creds *credentials.Credentials,
	options ...RequestOptions,
) (*http.Request, *http.Response, error) {
	// TODO handle context for clean exit
	// default options
	method := "GET"
	body := strings.NewReader("")

	// apply options if specified
	if len(options) > 0 {
		method = options[0].Method
		body = strings.NewReader(options[0].Body)
	}

	// create request
	url := p.GetRequestURL(region, service)
	logger.LogDebug("url", url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	// create signer to generate auth signature header
	signer := v4.NewSigner(creds)
	// sign request
	if _, err := signer.Sign(req, body, service, region, time.Now()); err != nil {
		return nil, nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// log success to console
	if res.StatusCode == http.StatusOK {
		logger.LogSuccess(service, p.Permission)
	} else {
		logger.LogDenied(res.StatusCode, service, p.Permission)
	}

	return req, res, nil
}

type RequestOptions struct {
	Method string `json:"method"`
	Body   string `json:"body"`
}
