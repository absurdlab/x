package errorx

// Coalesce returns the first non-nil error.
func Coalesce(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
