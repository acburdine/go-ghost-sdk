package contentapi

import (
	"encoding/json"
	"fmt"

	"github.com/acburdine/go-ghost-sdk/models"
	"github.com/pkg/errors"
)

type tagResponse struct {
	Tags []*models.Tag `json:"tags"`
	Meta *Meta         `json:"meta"`
}

// TagResource represents actions that can be done on the Tag resource
type TagResource struct {
	client *Client
}

// Browse returns a list of tags
func (t *TagResource) Browse(opts *BrowseOptions) ([]*models.Tag, *Meta, error) {
	resp, err := t.client.makeRequest("tags", opts.convert())
	if err != nil {
		return nil, nil, err
	}

	var respData tagResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, nil, err
	}

	return respData.Tags, respData.Meta, nil
}

// BrowseAll iterates through all authors, calling the callback with the result of each request
func (t *TagResource) BrowseAll(opts *BrowseOptions, fn func([]*models.Tag, bool) bool) error {
	hasMore, cont := true, true

	for page := int64(1); hasMore && cont; page++ {
		copyOpts := opts.copy()
		copyOpts.Page = page
		tags, meta, err := t.Browse(copyOpts)
		if err != nil {
			return err
		}

		hasMore = meta.Pagination.Pages > page
		cont = fn(tags, !hasMore)
	}

	return nil
}

// Read reads a specific author
func (t *TagResource) Read(id string, opts *ReadOptions) (*models.Tag, error) {
	endpoint := fmt.Sprintf("tags/%s/", id)
	resp, err := t.client.makeRequest(endpoint, opts.convert())
	if err != nil {
		return nil, err
	}

	var respData tagResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	if len(respData.Tags) == 0 {
		return nil, errors.New(fmt.Sprintf("no tag found with id %s", id))
	}

	return respData.Tags[0], nil
}
