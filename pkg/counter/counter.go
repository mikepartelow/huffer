package counter

func Counts[T comparable](s []T) map[T]int {
	counts := make(map[T]int)

	for _, r := range s {
		counts[r] += 1
	}

	return counts
}
