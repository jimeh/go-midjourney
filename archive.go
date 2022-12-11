package midjourney

import (
	"context"
	"net/url"
	"time"
)

func (c *Client) ArchiveDay(
	ctx context.Context,
	date time.Time,
) (jobIDs []string, err error) {
	err = c.Get(ctx, "app/archive/day", url.Values{
		"day":   []string{date.Format("2")},
		"month": []string{date.Format("1")},
		"year":  []string{date.Format("2006")},
	}, &jobIDs)

	return jobIDs, err
}
