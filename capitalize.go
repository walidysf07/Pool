package piscine

func Capitalize(s string) string {
	start := true
	result := ""

	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if start {
				if c >= 'a' && c <= 'z' {
					result += string(c - 32)
				} else {
					result += string(c)
				}
				start = false
			} else {
				if c >= 'A' && c <= 'Z' {
					result += string(c + 32)
				} else {
					result += string(c)
				}
			}
		} else {
			result += string(c)
			start = true
		}
	}
	return result
}
