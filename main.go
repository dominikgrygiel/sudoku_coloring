package main

import (
	"fmt"
)

func readSudoku() *Sudoku {
	var squareSize int
	fmt.Scanf("%d", &squareSize)

	sudoku := NewSudoku(squareSize)

	for i := 0; i < sudoku.NumColors; i++ {
		for j := 0; j < sudoku.NumColors; j++ {
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
