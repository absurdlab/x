package boolx

// Coalesce returns the first non-nil *bool.
func Coalesce(elem ...*bool) *bool {
	for _, each := range elem {
		if each != nil {
			return each
		}
	}
	return nil
}
