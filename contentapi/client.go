package contentapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// Client represents a Ghost Content API client
type Client struct {
	baseURL               *url.URL
	key, version, apiPath string
	httpClient            *http.Client

	Posts *PostResource
	Pages *PageResource
}

func (c *Client) makeRequest(endpoint string, params url.Values) (*http.Response, error) {
	params.Add("key", c.key)
	fullPath := strings.Join([]string{c.apiPath, "api", c.version, "content", endpoint}, "/")

	resourceURI, err := url.Parse(fullPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing endpoint uri %s", endpoint)
	}

	resourceURI.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", c.baseURL.ResolveReference(resourceURI).String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request object")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error making request")
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		var errResp errorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding api error response")
		}

		return resp, errResp.Errors
	}

	return resp, nil
}

// New takes in options and returns a new Client
func New(httpClient *http.Client, opts *ClientOptions) (*Client, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(opts.Host)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing host URL")
	}

	client := Client{
		baseURL:    baseURL,
		key:        opts.Key,
		version:    opts.Version,
		apiPath:    opts.GhostPath,
		httpClient: httpClient,
	}

	client.Posts = &PostResource{&client}
	client.Pages = &PageResource{&client}

	return &client, nil
}
