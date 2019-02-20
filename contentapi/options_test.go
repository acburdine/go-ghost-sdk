package contentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testHost    = "https://demo.ghost.io"
	testVersion = "v2"
	testKey     = "1234567890abcdef1234567890"
)

func TestClientOptions_ValidateMissingHost(t *testing.T) {
	opts := ClientOptions{
		Version: testVersion,
		Key:     testKey,
	}

	err := opts.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Host must be specified")
}

func TestClientOptions_ValidateMissingVersion(t *testing.T) {
	opts := ClientOptions{
		Host: testHost,
		Key:  testKey,
	}

	err := opts.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Version must be specified")
}

func TestClientOptions_ValidateUnsupportedVersion(t *testing.T) {
	opts := ClientOptions{
		Host:    testHost,
		Version: "v5",
		Key:     testKey,
	}

	err := opts.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Version v5 is not supported")
}

func TestClientOptions_ValidateMissingKey(t *testing.T) {
	opts := ClientOptions{
		Host:    testHost,
		Version: testVersion,
	}

	err := opts.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Key must be specified")
}

func TestClientOptions_ValidateKeyPattern(t *testing.T) {
	opts := ClientOptions{
		Host:    testHost,
		Version: testVersion,
		Key:     "invalid key",
	}

	err := opts.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Key must be 26 hex characters long")
}

func TestValidate_GhostPath(t *testing.T) {
	opts := ClientOptions{
		Host:    testHost,
		Version: testVersion,
		Key:     testKey,
	}

	err := opts.Validate()
	assert.Nil(t, err)
	assert.Equal(t, "ghost", opts.GhostPath)
}

func TestValidate_CustomGhostPath(t *testing.T) {
	opts := ClientOptions{
		Host:      testHost,
		Version:   testVersion,
		Key:       testKey,
		GhostPath: "customghost",
	}

	err := opts.Validate()
	assert.Nil(t, err)
	assert.Equal(t, "customghost", opts.GhostPath)
}
