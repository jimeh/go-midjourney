package midjourney

import (
	"net/url"
	"strings"

	"github.com/rs/zerolog"
)

type Option interface {
	apply(*APIClient) error
}

type optionFunc func(*APIClient) error

func (fn optionFunc) apply(o *APIClient) error {
	return fn(o)
}

// WithAuthToken returns a new Option type which sets the auth token that the
// client will use. The authToken value can be fetched from the
// "__Secure-next-auth.session-token" cookie on the midjourney.com website.
func WithAuthToken(authToken string) Option {
	return optionFunc(func(c *APIClient) error {
		c.AuthToken = authToken

		return nil
	})
}

func WithAPIURL(baseURL string) Option {
	return optionFunc(func(c *APIClient) error {
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

func WithHTTPClient(httpClient HTTPClient) Option {
	return optionFunc(func(c *APIClient) error {
		c.HTTPClient = httpClient

		return nil
	})
}

func WithUserAgent(userAgent string) Option {
	return optionFunc(func(c *APIClient) error {
		c.UserAgent = userAgent

		return nil
	})
}

func WithLogger(logger zerolog.Logger) Option {
	return optionFunc(func(c *APIClient) error {
		c.Logger = logger

		return nil
	})
}
