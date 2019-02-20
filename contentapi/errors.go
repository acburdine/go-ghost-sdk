package contentapi

import (
	"fmt"
	"strings"
)

// GhostError represents an error returned from the Ghost API
type GhostError struct {
	Message   string `json:"message"`
	ErrorType string `json:"errorType"`
}

// ErrorList represents a list of errors
type ErrorList []*GhostError

func (l ErrorList) Error() string {
	strs := make([]string, len(l))

	for _, e := range l {
		strs = append(strs, fmt.Sprintf("%s: %s", e.ErrorType, e.Message))
	}

	return fmt.Sprintf("%d error(s) occurred: %s", len(l), strings.Join(strs, ", "))
}

type errorResponse struct {
	Errors ErrorList `json:"errors"`
}
