package stringz

import "strings"

// SpaceJoin joins values with space
func SpaceJoin(values ...string) string {
	return strings.Join(values, " ")
}
