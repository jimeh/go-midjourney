package midjourney

import (
	"context"
	"fmt"
	"net/http"
)

var ErrJobIDsRequired = fmt.Errorf("%w: job IDs required", Err)

type collectionJobsRequest struct {
	CollectionID string   `json:"collection_id,omitempty"`
	JobIDs       []string `json:"job_ids,omitempty"`
}

type CollectionJobsResult struct {
	Failures  []string `json:"failures,omitempty"`
	Success   bool     `json:"success,omitempty"`
	Successes []string `json:"successes,omitempty"`
}

func (c *Client) CollectionJobsAdd(
	ctx context.Context,
	collectionID string,
	jobIDs []string,
) (*CollectionJobsResult, error) {
	return c.collectionJobs(ctx, http.MethodPut, collectionID, jobIDs)
}

func (c *Client) CollectionJobsRemove(
	ctx context.Context,
	collectionID string,
	jobIDs []string,
) (*CollectionJobsResult, error) {
	return c.collectionJobs(ctx, http.MethodDelete, collectionID, jobIDs)
}

func (c *Client) collectionJobs(
	ctx context.Context,
	method string,
	collectionID string,
	jobIDs []string,
) (*CollectionJobsResult, error) {
	if collectionID == "" {
		return nil, ErrCollectionIDRequired
	}
	if len(jobIDs) == 0 {
		return nil, ErrJobIDsRequired
	}

	var resp *CollectionJobsResult

	err := c.Request(
		ctx, method, "app/collections-jobs/", nil,
		&collectionJobsRequest{CollectionID: collectionID, JobIDs: jobIDs},
		resp,
	)

	return resp, err
}
