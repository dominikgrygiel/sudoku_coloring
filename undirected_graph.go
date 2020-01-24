package main

type Node struct {
	vertices map[*Node]bool
}

type UndirectedGraph struct {
	nodes []Node
}

func NewUndirectedGraph(numNodes int) *UndirectedGraph {
	graph := UndirectedGraph{make([]Node, numNodes)}

	for idx := range graph.nodes {
		graph.nodes[idx].vertices = make(map[*Node]bool)
	}

	return &graph
}

func (g *UndirectedGraph) AddEdge(n1, n2 *Node) {
	n1.vertices[n2] = true
	n2.vertices[n1] = true
}

func (n *Node) Vertices() (nodes []*Node) {
	for node, v := range n.vertices {
		if v == true {
			nodes = append(nodes, node)
		}
	}

	return
}
