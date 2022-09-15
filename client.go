package midjourney

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog"
)

var (
	Err                  = errors.New("midjourney")
	ErrNoAuthToken       = fmt.Errorf("%w: no auth token", Err)
	ErrInvalidAPIURL     = fmt.Errorf("%w: invalid API URL", Err)
	ErrInvalidHTTPClient = fmt.Errorf("%w: invalid HTTP client", Err)
	ErrResponseStatus    = fmt.Errorf("%w: response status", Err)

	DefaultAPIURL = url.URL{
		Scheme: "https",
		Host:   "www.midjourney.com",
		Path:   "/api/",
	}
	DefaultUserAgent = "go-midjourney/0.0.0-dev"
)

type Option interface {
	apply(*Client) error
}

type optionFunc func(*Client) error

func (fn optionFunc) apply(o *Client) error {
	return fn(o)
}

func WithAuthToken(authToken string) Option {
	return optionFunc(func(c *Client) error {
		c.AuthToken = authToken

		return nil
	})
}

func WithAPIURL(baseURL string) Option {
	return optionFunc(func(c *Client) error {
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}

		u, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.APIURL = u

		return nil
	})
}

func WithHTTPClient(httpClient *http.Client) Option {
	return optionFunc(func(c *Client) error {
		c.HTTPClient = httpClient

		return nil
	})
}

func WithUserAgent(userAgent string) Option {
	return optionFunc(func(c *Client) error {
		c.UserAgent = userAgent

		return nil
	})
}

func WithLogger(logger zerolog.Logger) Option {
	return optionFunc(func(c *Client) error {
		c.Logger = logger

		return nil
	})
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HTTPClient HTTPClient
	APIURL     *url.URL
	AuthToken  string
	UserAgent  string
	Logger     zerolog.Logger
}

func New(options ...Option) (*Client, error) {
	c := &Client{
		HTTPClient: http.DefaultClient,
		APIURL:     &DefaultAPIURL,
		UserAgent:  DefaultUserAgent,
		Logger:     zerolog.Nop(),
	}
	err := c.Set(options...)

	return c, err
}

func (c *Client) Set(options ...Option) error {
	for _, opt := range options {
		err := opt.apply(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.URL = c.APIURL.ResolveReference(req.URL)
	c.Logger.Debug().Str("url", req.URL.String()).Msg("request")

	req.Header.Set("Accept", "application/json")
	if c.AuthToken != "" {
		req.Header.Set(
			"Cookie", "__Secure-next-auth.session-token="+c.AuthToken,
		)
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return c.HTTPClient.Do(req)
}
