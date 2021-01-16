package stringz

import (
	"net/url"
	"strings"
)

// SpaceJoin joins values with space
func SpaceJoin(values ...string) string {
	return strings.Join(values, " ")
}

// IsURL checks if the value is a valid url.
func IsURL(v string) bool {
	_, err := url.ParseRequestURI(v)
	return err == nil
}
