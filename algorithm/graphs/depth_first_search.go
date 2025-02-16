package graphs

// represents a graph using an adjacency list
type GraphDFS struct {
	vertices map[int][]int
}

// creates a new graph
func NewGraphDFS() *GraphDFS {
	return &GraphDFS{
		vertices: make(map[int][]int),
	}
}

// adds an edge to the graph
func (g *GraphDFS) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
}

func (g *GraphDFS) DFSRecursive(start int, visited map[int]bool, result *[]int) {
	visited[start] = true
	*result = append(*result, start)

	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFSRecursive(neighbor, visited, result)
		}
	}
}

func (g *GraphDFS) DFSIterative(start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	result := []int{}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)

			// Add neighbors to the stack in reverse order for correct traversal
			for i := len(g.vertices[vertex]) - 1; i >= 0; i-- {
				neighbor := g.vertices[vertex][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	return result
}
