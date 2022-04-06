package cartesian

// The Algorithm is from https://rosettacode.org/wiki/Cartesian_product_of_two_or_more_lists#Go
func CartN(a ...[]string) (c [][]string) {
	if len(a) == 0 {
		return [][]string{nil}
	}
	r := CartN(a[1:]...)
	for _, e := range a[0] {
		for _, p := range r {
			c = append(c, append([]string{e}, p...))
		}
	}
	return
}
