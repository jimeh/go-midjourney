package midjourney

import (
	"context"
	"fmt"
	"net/url"
)

var (
	ErrCollectionIDRequired = fmt.Errorf("%w: collection id required", Err)
	ErrCollectionNotFound   = fmt.Errorf("%w: collection", ErrNotFound)
)

type Collection struct {
	CoverJobID         string          `json:"cover_job_id,omitempty"`
	Created            string          `json:"created,omitempty"`
	CreatorAvatarJobID string          `json:"creator_avatar_job_id,omitempty"`
	CreatorCoverJobID  string          `json:"creator_cover_job_id,omitempty"`
	CreatorID          string          `json:"creator_id,omitempty"`
	CreatorUsername    string          `json:"creator_username,omitempty"`
	Data               *CollectionData `json:"data,omitempty"`
	Description        string          `json:"description,omitempty"`
	Hidden             bool            `json:"hidden,omitempty"`
	ID                 string          `json:"id,omitempty"`
	NumJobs            int             `json:"num_jobs,omitempty"`
	Public             bool            `json:"public,omitempty"`
	PublicEditable     bool            `json:"public_editable,omitempty"`
	SearchTerms        []string        `json:"search_terms,omitempty"`
	Title              string          `json:"title,omitempty"`
	Workspaces         []string        `json:"workspaces,omitempty"`
}

type CollectionsQuery struct {
	UserID       string `url:"user_id,omitempty"`
	CollectionID string `url:"collection_id,omitempty"`
}

func (cq *CollectionsQuery) URLValues() url.Values {
	v := url.Values{}

	if cq.UserID != "" {
		v.Set("user_id", cq.UserID)
	}
	if cq.CollectionID != "" {
		v.Set("collection_id", cq.CollectionID)
	}

	return v
}

func (c *Client) Collections(
	ctx context.Context,
	query *CollectionsQuery,
) ([]*Collection, error) {
	var collections []*Collection

	err := c.API.Get(ctx, "app/collections/", query.URLValues(), &collections)

	return collections, err
}

func (c *Client) GetCollection(
	ctx context.Context,
	collectionID string,
) (*Collection, error) {
	if collectionID == "" {
		return nil, ErrCollectionIDRequired
	}

	q := &CollectionsQuery{CollectionID: collectionID}

	var cols []*Collection

	// Deletion of a collection is strangely done by setting the hidden flag to
	// true. This is a bit confusing, but it's how the API works.
	err := c.API.Get(ctx, "app/collections/", q.URLValues(), &cols)

	if len(cols) == 0 {
		return nil, fmt.Errorf("%w: id=%s", ErrCollectionNotFound, collectionID)
	}

	return cols[0], err
}

func (c *Client) PutCollection(
	ctx context.Context,
	collection *Collection,
) (*Collection, error) {
	var col *Collection

	err := c.API.Put(ctx, "app/collections/", nil, collection, col)

	return col, err
}

func (c *Client) DeleteCollection(
	ctx context.Context,
	collectionID string,
) (*Collection, error) {
	if collectionID == "" {
		return nil, ErrCollectionIDRequired
	}

	var col *Collection

	// Deletion of a collection is strangely done by setting the hidden flag to
	// true. This is a bit confusing, but it's how the API works.
	err := c.API.Put(
		ctx, "app/collections/", nil,
		&Collection{ID: collectionID, Hidden: true}, col,
	)

	return col, err
}
