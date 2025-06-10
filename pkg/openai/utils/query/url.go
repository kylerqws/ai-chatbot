package query

import (
	"net/url"
	"strconv"
)

// UrlQuery builds a URL-encoded query string with optional parameters.
type UrlQuery struct {
	params *url.Values
}

// NewUrlQuery creates a new UrlQuery instance.
func NewUrlQuery() *UrlQuery {
	return &UrlQuery{params: &url.Values{}}
}

// SetQueryStringParam adds a string parameter if the value is non-nil and not empty.
func (q *UrlQuery) SetQueryStringParam(key string, value *string) {
	if value != nil && *value != "" {
		q.SetParam(key, *value)
	}
}

// SetQueryUint8Param adds an unsigned integer parameter if value is non-nil and greater than zero.
func (q *UrlQuery) SetQueryUint8Param(key string, value *uint8) {
	if value != nil && *value > 0 {
		q.SetParam(key, strconv.FormatUint(uint64(*value), 10))
	}
}

// SetParam sets a query parameter with the given key and value.
func (q *UrlQuery) SetParam(key, value string) {
	q.params.Set(key, value)
}

// Encode encodes the query parameters as a URL query string.
// Returns an empty string if no parameters are set.
func (q *UrlQuery) Encode() string {
	if str := q.params.Encode(); str != "" {
		return "?" + str
	}
	return ""
}
