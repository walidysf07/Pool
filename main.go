package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	for _, s := range a {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
