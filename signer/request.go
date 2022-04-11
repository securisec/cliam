package signer

import (
	"context"
	"net/http"
	"net/url"
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
) (*http.Request, *http.Response, error) {
	// TODO handle context for clean exit
	// default options
	method := "GET"
	body := strings.NewReader("")

	if p.Method != "" {
		method = p.Method
	}

	// get request requestURL
	requestURL := p.GetRequestURL(region, service)
	logger.LogDebug("url", requestURL)

	// if form data, set body
	if p.FormData != nil {
		form := url.Values{}
		for k, v := range p.FormData {
			form.Add(k, v)
		}
		body = strings.NewReader(form.Encode())
	}

	// create request
	req, err := http.NewRequest(method, requestURL, body)
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
