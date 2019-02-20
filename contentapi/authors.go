package contentapi

import (
	"encoding/json"
	"fmt"

	"github.com/acburdine/go-ghost-sdk/models"
	"github.com/pkg/errors"
)

type authorResponse struct {
	Authors []*models.Author `json:"authors"`
	Meta    *Meta            `json:"meta"`
}

// AuthorResource represents actions that can be done on the Author resource
type AuthorResource struct {
	client *Client
}

// Browse returns a list of authors
func (a *AuthorResource) Browse(opts *BrowseOptions) ([]*models.Author, *Meta, error) {
	resp, err := a.client.makeRequest("authors", opts.convert())
	if err != nil {
		return nil, nil, err
	}

	var respData authorResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, nil, err
	}

	return respData.Authors, respData.Meta, nil
}

// BrowseAll iterates through all authors, calling the callback with the result of each request
func (a *AuthorResource) BrowseAll(opts *BrowseOptions, fn func([]*models.Author, bool) bool) error {
	hasMore, cont := true, true

	for page := int64(1); hasMore && cont; page++ {
		copyOpts := opts.copy()
		copyOpts.Page = page
		authors, meta, err := a.Browse(copyOpts)
		if err != nil {
			return err
		}

		hasMore = meta.Pagination.Pages > page
		cont = fn(authors, !hasMore)
	}

	return nil
}

// Read reads a specific author
func (a *AuthorResource) Read(id string, opts *ReadOptions) (*models.Author, error) {
	endpoint := fmt.Sprintf("authors/%s/", id)
	resp, err := a.client.makeRequest(endpoint, opts.convert())
	if err != nil {
		return nil, err
	}

	var respData authorResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	if len(respData.Authors) == 0 {
		return nil, errors.New(fmt.Sprintf("no author found with id %s", id))
	}

	return respData.Authors[0], nil
}
