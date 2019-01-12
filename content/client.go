package content

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/pkg/errors"
)

var (
	supportedVersions = []string{"v2"}
	keyRegex          = regexp.MustCompile("[0-9a-f]{26}")
)

type Options struct {
	Host      string
	GhostPath string
	Version   string
	Key       string
}

type Client struct {
	baseURL               *url.URL
	key, version, apiPath string
	httpClient            *http.Client
}

func validateOptions(opts *Options) error {
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
}

func New(opts *Options, httpClient *http.Client) (*Client, error) {
	if err := validateOptions(opts); err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(opt)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing host URL")
	}

	return nil, nil
}
