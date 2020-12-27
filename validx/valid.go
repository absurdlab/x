package validx

type (
	// ErrFunc is a function that emits an error
	ErrFunc func() error
	// StrFunc is a function that emits error based
	// on the string input
	StrFunc func(string) error
	// StrArrFunc is a function that emits error
	StrArrFunc func([]string) error
)

// Err returns an ErrFunc to return the given error.
func Err(err error) ErrFunc {
	return func() error {
		return err
	}
}

// AnyErr executes the list of errFunc in sequence, and returns
// the first error returned by one of the errFunc, if any.
func AnyErr(errFunc ...ErrFunc) error {
	for _, f := range errFunc {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// ValidStr returns an ErrFunc which invokes StrFunc with
// the string value.
func ValidStr(value string, f StrFunc) ErrFunc {
	return func() error {
		return f(value)
	}
}

// ValidStrElem returns an ErrFunc which invoke StrFunc on
// each element in the string slice values.
func ValidStrElem(values []string, f StrFunc) ErrFunc {
	return func() error {
		for _, v := range values {
			if err := f(v); err != nil {
				return err
			}
		}
		return nil
	}
}

// ValidStrArr returns an ErrFunc which invokes StrArrFunc with
// the string slice values.
func ValidStrArr(values []string, f StrArrFunc) ErrFunc {
	return func() error {
		return f(values)
	}
}

// NonEmptyStr returns an ErrFunc which checks if the string is not empty.
// If it is empty, errFunc is invoked to produce an error.
func NonEmptyStr(value string, errFunc ErrFunc) ErrFunc {
	return func() error {
		if len(value) > 0 {
			return nil
		}
		return errFunc()
	}
}

// NonEmptyStrArr returns an ErrFunc which checks if the string array
// is not empty. If it is empty, errFunc is invoked to produce an error.
func NonEmptyStrArr(values []string, errFunc ErrFunc) ErrFunc {
	return func() error {
		if len(values) > 0 {
			return nil
		}
		return errFunc()
	}
}

// OptionalStr returns an error func that invokes errFunc only when the given value
// is not empty.
func OptionalStr(value string, errFunc ErrFunc) ErrFunc {
	return func() error {
		if len(value) == 0 {
			return nil
		}
		return errFunc()
	}
}

// NonNil returns an error func that invokes errFunc only when the given value is nil.
func NonNil(value interface{}, errFunc ErrFunc) ErrFunc {
	return func() error {
		if value != nil {
			return nil
		}
		return errFunc()
	}
}
