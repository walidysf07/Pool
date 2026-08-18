package piscine

func LastRune(s string) rune {
	Str := []rune(s)

	if len(Str) == 0 {
		return 0
	}
	return Str[len(Str)-1]
}
