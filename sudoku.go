package main

import (
	"strconv"
)

type sudokuColors [][]int

type Sudoku struct {
	board      *sudokuBoard
	colors     sudokuColors
	squareSize int
	NumColors  int
}

func NewSudoku(squareSize int) *Sudoku {
	numColors := squareSize * squareSize

	colors := make(sudokuColors, numColors)
	for i := range colors {
		colors[i] = make([]int, numColors)
	}

	board := newSudokuBoard(squareSize)
	sudoku := Sudoku{board, colors, squareSize, numColors}

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
	for i := 0; i < s.NumColors; i++ {
		for j := 0; j < s.NumColors; j++ {
			if s.GetColor(i, j) < 1 {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) duplicateColors() (originalColors sudokuColors) {
	originalColors = s.colors

	colorsCopy := make(sudokuColors, s.NumColors)
	for i := range colorsCopy {
		colorsCopy[i] = make([]int, s.NumColors)
		copy(colorsCopy[i], s.colors[i])
	}

	s.colors = colorsCopy

	return
}

func colorDifference(colors []int, numColors int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range colors {
		m[item] = true
	}
	for i := 1; i <= numColors; i++ {
		if _, ok := m[i]; !ok {
			diff = append(diff, i)
		}
	}

	return
}

func doSolve(s Sudoku) (bool, *sudokuColors) {
	var minAmbigiousColors []int
	var minAmbigiousCoords [2]int

	for assignedColor := true; assignedColor; assignedColor = false {
		minAmbigiousColors = []int{}

		for i := 0; i < s.NumColors; i++ {
			for j := 0; j < s.NumColors; j++ {
				if s.GetColor(i, j) < 1 {
					var usedColors []int
					neighbours := s.board.nodeBoard[i][j].Vertices()

					for idx := range neighbours {
						if color := s.GetNodeColor(neighbours[idx]); color > 0 {
							usedColors = append(usedColors, color)
						}
					}

					possibleColors := colorDifference(usedColors, s.NumColors)
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
			colorsOrig := s.duplicateColors()
			s.SetColor(minAmbigiousCoords[0], minAmbigiousCoords[1], minAmbigiousColors[colorIdx])

			if solved, colors := doSolve(s); solved {
				return true, colors
			}

			s.colors = colorsOrig
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
	for i := 0; i < s.NumColors; i++ {
		for j := 0; j < s.NumColors; j++ {
			repr += strconv.Itoa(s.GetColor(i, j))

			if j < s.NumColors-1 {
				repr += " "
			}
		}

		if i < s.NumColors-1 {
			repr += "\n"
		}
	}

	return
}
