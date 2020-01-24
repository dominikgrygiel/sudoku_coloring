package main

import (
	"fmt"
)

func readSudoku() *Sudoku {
	sudoku := NewSudoku()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var x int
			fmt.Scanf("%d", &x)

			sudoku.SetColor(i, j, x)
		}
	}

	return sudoku
}

func main() {
	sudoku := readSudoku()
	sudoku.Solve()
	fmt.Println(sudoku)
}
