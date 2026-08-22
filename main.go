package main

import (
	"fmt"
	"os"
)

func valid(board [9][9]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
	}

	startRow := row / 3 * 3
	startCol := col / 3 * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}

	return true
}

func solve(board *[9][9]int, solution *[9][9]int, solutions *int) {
	if *solutions > 1 {
		return
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				continue
			}

			for num := 1; num <= 9; num++ {
				if valid(*board, row, col, num) {
					board[row][col] = num

					solve(board, solution, solutions)

					board[row][col] = 0

					if *solutions > 1 {
						return
					}
				}
			}

			return
		}
	}

	*solutions++
	*solution=*board 
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var board [9][9]int

	for row := 0; row < 9; row++ {
		input := os.Args[row+1]

		if len(input) != 9 {
			fmt.Println("Error")
			return
		}

		for col := 0; col < 9; col++ {
			c := input[col]

			if c == '.' {
				continue
			}

			if c < '1' || c > '9' {
				fmt.Println("Error")
				return
			}

			num := int(c - '0')

			if !valid(board, row, col, num) {
				fmt.Println("Error")
				return
			}

			board[row][col] = num
		}
	}

	var solution [9][9]int 
	solutions := 0
	solve(&board, &solution, &solutions)

	if solutions != 1 {
		fmt.Println("Error")
		return
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col > 0 {
				fmt.Print(" ")
			}
			fmt.Print(solution[row][col])
		}
		fmt.Println()
	}
}
