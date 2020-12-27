package stringsx

// Ref returns the reference to a new string whose value equals the given string.
func Ref(value string) *string {
	return &value
}
