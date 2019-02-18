package contentapi

import (
	"encoding/json"
	"net/url"

	"github.com/acburdine/go-ghost-sdk/models"
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
// TODO(acb): support passing in qs options
func (p *PostResource) Browse() ([]*models.Post, *Meta, error) {
	req, err := p.client.buildRequest("posts", url.Values{})
	if err != nil {
		return nil, nil, err
	}

	// TODO(acb): centralized "Do" method
	resp, err := p.client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var respData postResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, nil, err
	}

	return respData.Posts, respData.Meta, nil
}
