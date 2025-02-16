package graphs

type GraphBFS struct {
	vertices map[int][]int
}

func NewGraphBFS() *GraphBFS {
	return &GraphBFS{
		vertices: make(map[int][]int),
	}
}

// adds an edge to the graph
func (g *GraphBFS) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
}

// BFS performs a breadth-first search
func (g *GraphBFS) BFS(start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, neighbor := range g.vertices[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}
