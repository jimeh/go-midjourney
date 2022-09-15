package midjourney

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) ArchiveDay(
	ctx context.Context,
	date time.Time,
) (jobIDs []string, err error) {
	u := &url.URL{
		Path: "app/archive/day/",
		RawQuery: url.Values{
			"day":   []string{date.Format("2")},
			"month": []string{date.Format("1")},
			"year":  []string{date.Format("2006")},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrResponseStatus, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&jobIDs)
	if err != nil {
		return nil, err
	}

	return jobIDs, nil
}
