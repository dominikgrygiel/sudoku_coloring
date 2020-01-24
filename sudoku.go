package main

import (
	"strconv"
)

type sudokuColors [9][9]int

type Sudoku struct {
	board  *sudokuBoard
	colors sudokuColors
}

func NewSudoku() *Sudoku {
	var colors sudokuColors
	board := newSudokuBoard()
	sudoku := Sudoku{board, colors}

	return &sudoku
}

func (s *Sudoku) GetNodeColor(n *Node) int {
	coord, _ := s.board.nodeLookup[n]
	return s.GetColor(coord[0], coord[1])
}

func (s *Sudoku) GetColor(i, j int) int {
	return s.colors[i][j]
}

func (s *Sudoku) SetColor(i, j, color int) {
	s.colors[i][j] = color
}

func (s *Sudoku) IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.colors[i][j] < 1 {
				return false
			}
		}
	}

	return true
}

func colorDifference(colors []int) (diff []int) {
	availableColors := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := make(map[int]bool)

	for _, item := range colors {
		m[item] = true
	}
	for _, item := range availableColors {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return
}

func doSolve(s Sudoku) (bool, *sudokuColors) {
	var minAmbigiousColors []int
	var minAmbigiousCoords [2]int

	for assignedColor := true; assignedColor; assignedColor = false {
		minAmbigiousColors = []int{}

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if s.GetColor(i, j) < 1 {
					var usedColors []int
					neighbours := s.board.nodeBoard[i][j].Vertices()

					for idx := range neighbours {
						if color := s.GetNodeColor(neighbours[idx]); color > 0 {
							usedColors = append(usedColors, color)
						}
					}

					possibleColors := colorDifference(usedColors)
					if numPossibleColors := len(possibleColors); numPossibleColors == 1 {
						s.SetColor(i, j, possibleColors[0])
						assignedColor = true
					} else if numPossibleColors > 1 && (len(minAmbigiousColors) == 0 || numPossibleColors < len(minAmbigiousColors)) {
						minAmbigiousColors = possibleColors
						minAmbigiousCoords = [2]int{i, j}
					}
				}
			}
		}
	}

	if !s.IsSolved() {
		for colorIdx := range minAmbigiousColors {
			s.SetColor(minAmbigiousCoords[0], minAmbigiousCoords[1], minAmbigiousColors[colorIdx])

			if solved, colors := doSolve(s); solved {
				return true, colors
			}
		}

		return false, &s.colors
	}

	return true, &s.colors
}

func (s *Sudoku) Solve() {
	_, colors := doSolve(*s)
	s.colors = *colors
}

func (s *Sudoku) String() (repr string) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			repr += strconv.Itoa(s.colors[i][j])

			if j < 8 {
				repr += " "
			}
		}

		if i < 8 {
			repr += "\n"
		}
	}

	return
}
