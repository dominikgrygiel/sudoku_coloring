package main

type sudokuNodeBoard [][]*Node
type sudokuNodeLookup map[*Node][2]int
type sudokuBoard struct {
	nodeBoard  sudokuNodeBoard
	nodeLookup sudokuNodeLookup
	graph      *UndirectedGraph
	squareSize int
	numColors  int
}

func (s *sudokuBoard) linkRows() {
	for i := 0; i < s.numColors; i++ {
		for j := 0; j < s.numColors; j++ {
			for k := j + 1; k < s.numColors; k++ {
				s.graph.AddEdge(s.nodeBoard[i][j], s.nodeBoard[i][k])
			}
		}
	}
}

func (s *sudokuBoard) linkColumns() {
	for i := 0; i < s.numColors; i++ {
		for j := 0; j < s.numColors; j++ {
			for k := j + 1; k < s.numColors; k++ {
				s.graph.AddEdge(s.nodeBoard[j][i], s.nodeBoard[k][i])
			}
		}
	}
}

func (s *sudokuBoard) linkSquares() {
	for i := 0; i < s.numColors; i += s.squareSize {
		for j := 0; j < s.numColors; j += s.squareSize {
			for k := 0; k < s.squareSize; k++ {
				for l := 0; l < s.squareSize; l++ {
					for m := 0; m < s.squareSize; m++ {
						for n := 0; n < s.squareSize; n++ {
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

func newSudokuBoard(squareSize int) *sudokuBoard {
	numColors := squareSize * squareSize
	graph := NewUndirectedGraph(numColors * numColors)

	nodeLookup := make(sudokuNodeLookup)
	nodeBoard := make(sudokuNodeBoard, numColors)
	for i := range nodeBoard {
		nodeBoard[i] = make([]*Node, numColors)
	}

	for i := 0; i < numColors; i++ {
		for j := 0; j < numColors; j++ {
			node := &graph.nodes[numColors*i+j]
			nodeBoard[i][j] = node
			nodeLookup[node] = [2]int{i, j}
		}
	}

	board := sudokuBoard{nodeBoard, nodeLookup, graph, squareSize, numColors}
	board.linkRows()
	board.linkColumns()
	board.linkSquares()

	return &board
}
