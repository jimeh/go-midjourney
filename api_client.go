package midjourney

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
	HTTPClient HTTPClient
	APIURL     *url.URL
	AuthToken  string
	UserAgent  string
	Logger     zerolog.Logger
}

func NewAPI(options ...Option) (*APIClient, error) {
	c := &APIClient{
		HTTPClient: http.DefaultClient,
		APIURL:     &DefaultAPIURL,
		UserAgent:  DefaultUserAgent,
		Logger:     zerolog.Nop(),
	}
	err := c.Set(options...)

	return c, err
}

func (ac *APIClient) Set(options ...Option) error {
	for _, opt := range options {
		err := opt.apply(ac)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ac *APIClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	if ac.AuthToken != "" {
		req.Header.Set(
			"Cookie", "__Secure-next-auth.session-token="+ac.AuthToken,
		)
	}
	if ac.UserAgent != "" {
		req.Header.Set("User-Agent", ac.UserAgent)
	}

	return ac.HTTPClient.Do(req)
}

func (ac *APIClient) Request(
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
	u = ac.APIURL.ResolveReference(u)

	ac.Logger.Debug().
		Str("method", method).
		Str("url", u.String()).
		Msg("request")

	var req *http.Request
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}

		ac.Logger.Trace().RawJSON("body", b).Msg("request")

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

	resp, err := ac.Do(req)
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

	ac.Logger.Trace().RawJSON("body", buf.Bytes()).Msg("response")

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

func (ac *APIClient) Get(
	ctx context.Context,
	path string,
	params url.Values,
	x any,
) error {
	return ac.Request(ctx, http.MethodGet, path, params, nil, x)
}

func (ac *APIClient) Put(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return ac.Request(ctx, http.MethodPut, path, params, body, x)
}

func (ac *APIClient) Post(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return ac.Request(ctx, http.MethodPost, path, params, body, x)
}

func (ac *APIClient) Patch(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return ac.Request(ctx, http.MethodPatch, path, params, body, x)
}

func (ac *APIClient) Delete(
	ctx context.Context,
	path string,
	params url.Values,
	body any,
	x any,
) error {
	return ac.Request(ctx, http.MethodDelete, path, params, body, x)
}
