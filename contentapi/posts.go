package contentapi

import (
	"encoding/json"
	"fmt"

	"github.com/acburdine/go-ghost-sdk/models"
	"github.com/pkg/errors"
)

type postResponse struct {
	Posts []*models.Post `json:"posts"`
	Meta  *Meta          `json:"meta"`
}

// PostResource represents actions that can be done on the Posts resource
type PostResource struct {
	client *Client
}

// Browse returns a list of posts
func (p *PostResource) Browse(opts *BrowseOptions) ([]*models.Post, *Meta, error) {
	resp, err := p.client.makeRequest("posts", opts.convert())
	if err != nil {
		return nil, nil, err
	}

	var respData postResponse
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, nil, err
	}

	return respData.Posts, respData.Meta, nil
}

// BrowseAll iterates through all posts, calling the callback with the result of each request
func (p *PostResource) BrowseAll(opts *BrowseOptions, fn func([]*models.Post, bool) bool) error {
	hasMore, cont := true, true

	for page := int64(1); hasMore && cont; page++ {
		copyOpts := opts.copy()
		copyOpts.Page = page
		posts, meta, err := p.Browse(copyOpts)
		if err != nil {
			return err
		}

		hasMore = meta.Pagination.Pages > page
		cont = fn(posts, !hasMore)
	}

	return nil
}

func (p *PostResource) Read(id string, opts *ReadOptions) (*models.Post, error) {
	endpoint := fmt.Sprintf("posts/%s/", id)
	resp, err := p.client.makeRequest(endpoint, opts.convert())
	if err != nil {
		return nil, err
	}

	var respData postResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	if len(respData.Posts) == 0 {
		return nil, errors.New(fmt.Sprintf("no post found with id %s", id))
	}

	return respData.Posts[0], nil
}
