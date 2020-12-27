package boolx

var (
	t = true
	f = false
)

// TrueRef returns the reference to a true boolean
func TrueRef() *bool {
	return &t
}

// FalseRef returns the reference to a false boolean
func FalseRef() *bool {
	return &f
}

// Ref returns the reference to the value of the given bool.
func Ref(value bool) *bool {
	return &value
}
