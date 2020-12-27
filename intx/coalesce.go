package intx

func FirstPositiveOrZero(elem ...int64) int64 {
	for _, each := range elem {
		if each > 0 {
			return each
		}
	}
	return 0
}
