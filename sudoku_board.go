package main

type sudokuNodeBoard [9][9]*Node
type sudokuNodeLookup map[*Node][2]int
type sudokuBoard struct {
	nodeBoard  sudokuNodeBoard
	nodeLookup sudokuNodeLookup
	graph      *UndirectedGraph
}

func (s *sudokuBoard) linkRows() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := j + 1; k < 9; k++ {
				s.graph.AddEdge(s.nodeBoard[i][j], s.nodeBoard[i][k])
			}
		}
	}
}

func (s *sudokuBoard) linkColumns() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := j + 1; k < 9; k++ {
				s.graph.AddEdge(s.nodeBoard[j][i], s.nodeBoard[k][i])
			}
		}
	}
}

func (s *sudokuBoard) linkSquares() {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					for m := 0; m < 3; m++ {
						for n := 0; n < 3; n++ {
							n1 := s.nodeBoard[j+k][i+l]
							n2 := s.nodeBoard[j+m][i+n]

							if n1 != n2 {
								s.graph.AddEdge(n1, n2)
							}
						}
					}
				}
			}
		}
	}
}

func newSudokuBoard() *sudokuBoard {
	var nodeBoard sudokuNodeBoard
	nodeLookup := make(sudokuNodeLookup)
	graph := NewUndirectedGraph(9 * 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			node := &graph.nodes[9*i+j]
			nodeBoard[i][j] = node
			nodeLookup[node] = [2]int{i, j}
		}
	}

	board := sudokuBoard{nodeBoard, nodeLookup, graph}
	board.linkRows()
	board.linkColumns()
	board.linkSquares()

	return &board
}
