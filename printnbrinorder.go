package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	count := [10]int{}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		dgt := n % 10
		count[dgt]++
		n = n / 10
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < count[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
