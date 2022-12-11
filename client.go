package midjourney

import (
	"errors"
	"fmt"
	"net/url"
)

var (
	Err                  = errors.New("midjourney")
	ErrNoAuthToken       = fmt.Errorf("%w: no auth token", Err)
	ErrInvalidAuthToken  = fmt.Errorf("%w: invalid auth token", Err)
	ErrInvalidAPIURL     = fmt.Errorf("%w: invalid API URL", Err)
	ErrInvalidHTTPClient = fmt.Errorf("%w: invalid HTTP client", Err)
	ErrNotFound          = fmt.Errorf("%w: not found", Err)
	ErrResponse          = fmt.Errorf("%w: response", Err)
	ErrResponseStatus    = fmt.Errorf("%w: response status", ErrResponse)

	DefaultAPIURL = url.URL{
		Scheme: "https",
		Host:   "www.midjourney.com",
		Path:   "/api/",
	}

	DefaultUserAgent = "go-midjourney/0.0.1" // x-release-please-version
)

type Client struct {
	API *APIClient
}

func New(options ...Option) (*Client, error) {
	api, err := NewAPI(options...)
	if err != nil {
		return nil, err
	}

	return &Client{API: api}, nil
}

func (ac *Client) Set(options ...Option) error {
	return ac.API.Set(options...)
}
