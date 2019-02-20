package contentapi

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var (
	supportedVersions = []string{"v2"}
	keyRegex          = regexp.MustCompile("[0-9a-f]{26}")
)

// ClientOptions represents options given when creating a Client
type ClientOptions struct {
	Host      string
	GhostPath string
	Version   string
	Key       string
}

// Validate validations the options
func (opts *ClientOptions) Validate() error {
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

// ReadOptions represents options that can be passed to Read methods
type ReadOptions struct {
	Include string
	Fields  string
	Formats string
}

func (o *ReadOptions) convert() url.Values {
	values := url.Values{}

	if o.Include != "" {
		values.Add("include", o.Include)
	}
	if o.Fields != "" {
		values.Add("fields", o.Fields)
	}
	if o.Formats != "" {
		values.Add("formats", o.Formats)
	}

	return values
}

// BrowseOptions represents options that can be passed to Browse methods
type BrowseOptions struct {
	ReadOptions

	Filter string
	Limit  int64
	Page   int64
	Order  string
}

func (o *BrowseOptions) convert() url.Values {
	values := o.ReadOptions.convert()

	if o.Filter != "" {
		values.Add("filter", o.Filter)
	}
	if o.Limit > 0 {
		values.Add("limit", strconv.Itoa(int(o.Limit)))
	}
	if o.Page > 0 {
		values.Add("page", strconv.Itoa(int(o.Page)))
	}
	if o.Order != "" {
		values.Add("order", o.Order)
	}

	return values
}

func (o *BrowseOptions) copy() *BrowseOptions {
	return &BrowseOptions{
		ReadOptions: o.ReadOptions,
		Filter:      o.Filter,
		Limit:       o.Limit,
		Page:        o.Page,
		Order:       o.Order,
	}
}
