package contentapi

// Pagination represents pagination info returned from browse endpoints
type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Pages int64 `json:"pages"`
	Total int64 `json:"total"`
	Next  int64 `json:"next"`
	Prev  int64 `json:"prev"`
}

// Meta represents the meta information returned from browse endpoints
type Meta struct {
	Pagination *Pagination `json:"pagination"`
}
