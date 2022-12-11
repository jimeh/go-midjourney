package midjourney

import "context"

type CollectionData struct {
	Filters *CollectionFilters `json:"filters,omitempty"`
}

type CollectionFilters struct {
	OrderBy           string `json:"orderBy,omitempty"`
	JobType           string `json:"jobType,omitempty"`
	UserIDRankedScore string `json:"user_id_ranked_score,omitempty"`
	ShowFilters       bool   `json:"showFilters,omitempty"`
}

func (c *Client) PutCollectionData(
	ctx context.Context,
	collectionID string,
	data *CollectionData,
) (*Collection, error) {
	if collectionID == "" {
		return nil, ErrCollectionIDRequired
	}

	req := &Collection{
		ID:   collectionID,
		Data: data,
	}
	resp := &Collection{}

	err := c.API.Put(ctx, "app/collections/", nil, req, resp)

	return resp, err
}

func (c *Client) PutCollectionFilters(
	ctx context.Context,
	collectionID string,
	filters *CollectionFilters,
) (*Collection, error) {
	if collectionID == "" {
		return nil, ErrCollectionIDRequired
	}

	req := &Collection{
		ID: collectionID,
		Data: &CollectionData{
			Filters: filters,
		},
	}
	resp := &Collection{}

	err := c.API.Put(ctx, "app/collections/", nil, req, resp)

	return resp, err
}
