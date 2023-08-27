func FindMultiples(base int, limit int) []int {
	multiples := []int{}
	for i := base; i <= limit; i++ {
		if i%base == 0 {
			multiples = append(multiples, i)
		}
	}
	return multiples
}
