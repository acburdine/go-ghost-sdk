package contentapi

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	supportedVersions = []string{"v2"}
	keyRegex          = regexp.MustCompile("[0-9a-f]{26}")
)

// Options represents options given when creating a Client
type Options struct {
	Host      string
	GhostPath string
	Version   string
	Key       string
}

func (opts *Options) validate() error {
	if opts.Host == "" {
		return errors.New("error: Host must be specified")
	}
	if opts.Version == "" {
		return errors.New("error: Version must be specified")
	}
	if opts.Key == "" {
		return errors.New("error: Key must be specified")
	}
	if !keyRegex.MatchString(opts.Key) {
		return errors.New("error: Key must be 26 hex characters long")
	}

	isSupportedVersion := false
	for _, v := range supportedVersions {
		if v == opts.Version {
			isSupportedVersion = true
		}
	}
	if !isSupportedVersion {
		return errors.New(fmt.Sprintf("error: Version %s is not supported by the Content API", opts.Version))
	}

	if opts.GhostPath == "" {
		opts.GhostPath = "ghost"
	}

	return nil
}

// Client represents a Ghost Content API client
type Client struct {
	baseURL               *url.URL
	key, version, apiPath string
	httpClient            *http.Client

	Posts *PostEndpoints
}

func (c *Client) buildRequest(endpoint string, params url.Values) (*http.Request, error) {
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

	return req, nil
}

// New takes in options and returns a new Client
func New(httpClient *http.Client, opts *Options) (*Client, error) {
	if err := opts.validate(); err != nil {
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

	client.Posts = &PostEndpoints{&client}

	return &client, nil
}
