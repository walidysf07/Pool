package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)

	for i := 0; i < len(ra) && i < len(rb); i++ {
		if ra[i] < rb[i] {
			return -1
		} else if ra[i] > rb[i] {
			return 1
		}
	}
	if len(ra) < len(rb) {
		return -1
	} else if len(ra) > len(rb) {
		return 1
	}
	return 0
}
