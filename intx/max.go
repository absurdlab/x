package intx

// Max64 returns the maximum int64.
func Max64(elem ...int64) int64 {
	if len(elem) == 0 {
		return 0
	}

	max := elem[0]
	for _, i := range elem {
		if i > max {
			max = i
		}
	}

	return max
}
