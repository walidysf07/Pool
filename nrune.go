package piscine

func NRune(s string, n int) rune {
	Str := []rune(s)

	if n <= 0 || n > len(Str) {
		return 0
	}
	return Str[n-1]
}
