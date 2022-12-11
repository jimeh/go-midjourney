package midjourney

import (
	"bytes"
	"context"
	"encoding/json"
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
	DefaultUserAgent = "go-midjourney/0.0.0-dev"
)

type Option interface {
	apply(*Client) error
}

type optionFunc func(*Client) error

func (fn optionFunc) apply(o *Client) error {
	return fn(o)
}

// WithAuthToken returns a new Option type which sets the auth token that the
// client will use. The authToken value can be fetched from the
// "__Secure-next-auth.session-token" cookie on the midjourney.com website.
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

func (c *Client) Request(
	ctx context.Context,
	method string,
	path string,
	params url.Values,
	body any,
	result any,
) error {
	u := &url.URL{Path: path}
	if params != nil {
		u.RawQuery = params.Encode()
	}
	u = c.APIURL.ResolveReference(u)

	c.Logger.Debug().Str("method", method).Str("url", u.String()).Msg("request")

	var req *http.Request
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}

		c.Logger.Trace().RawJSON("body", b).Msg("request")

		buf := bytes.NewBuffer(b)
		req, err = http.NewRequestWithContext(ctx, method, u.String(), buf)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		var err error
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return err
		}
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s", ErrResponseStatus, resp.Status)
	}

	// When token is invalid, a HTTP 200 response with content type text/html is
	// returned. Hence we treat non-JSON responses as an invalid auth token
	// error.
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return ErrInvalidAuthToken
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	c.Logger.Trace().RawJSON("body", buf.Bytes()).Msg("response")

	err = json.Unmarshal(buf.Bytes(), result)
	if err != nil {
		respErr := &ResponseError{}
		unmarshalErr := json.Unmarshal(buf.Bytes(), respErr)
		if unmarshalErr != nil {
			return err
		}

		return respErr
	}

	return nil
}

func (c *Client) Get(
	ctx context.Context,
	path string,
	params url.Values,
	x any,
) error {
	return c.Request(ctx, http.MethodGet, path, params, nil, x)
}

func (c *Client) Put(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return c.Request(ctx, http.MethodPut, path, params, body, x)
}

func (c *Client) Post(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return c.Request(ctx, http.MethodPost, path, params, body, x)
}

func (c *Client) Patch(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return c.Request(ctx, http.MethodPatch, path, params, body, x)
}

func (c *Client) Delete(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return c.Request(ctx, http.MethodDelete, path, params, body, x)
}
