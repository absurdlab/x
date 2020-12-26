package strconvx

import "strconv"

// ParseInt64OrDefault parses the given value to int64, or returns the default value if the given value is empty
// or is otherwise malformed as an integer.
func ParseInt64OrDefault(value string, defaultValue int64) int64 {
	if len(value) == 0 {
		return defaultValue
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return i
}
