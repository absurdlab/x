package stringsx

import "strings"

// Delimited is a string containing zero or more tokens delimited by a common rune.
type Delimited string

// BySpace returns the tokens delimited by space in a slice.
func (d Delimited) BySpace() []string {
	return strings.Fields(d.String())
}

// ByComma returns the tokens delimited by comma in a slice.
func (d Delimited) ByComma() []string {
	return d.By(',')
}

// By returns the tokens delimited by the given rune in a slice.
func (d Delimited) By(dr rune) []string {
	return strings.FieldsFunc(d.String(), func(r rune) bool {
		return dr == r
	})
}

// String returns the original value of the Delimited
func (d Delimited) String() string {
	return string(d)
}
