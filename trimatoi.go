package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}

	sign := 1
	i := 0
	start := false
	result := 0

	for ; i < len(s); i++ {
		if s[i] == '-' && !start {
			sign = -1
		} else if s[i] == '+' && !start {
			sign = 1
		} else if s[i] >= '0' && s[i] <= '9' {
			start = true
			digit := int(s[i] - '0')
			result = result*10 + digit
		}
	}

	return result * sign
}
