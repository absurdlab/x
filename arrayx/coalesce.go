package arrayx

// Coalesce returns the first non-empty string array
func Coalesce(arrays ...[]string) []string {
	for _, each := range arrays {
		if len(each) > 0 {
			return each
		}
	}
	return []string{}
}
