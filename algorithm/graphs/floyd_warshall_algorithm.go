package graphs

import "math"

type GraphFloydWarshall struct {
	Graph [][]int
}

func FloydWarshall(g GraphFloydWarshall) [][]int {
	n := len(g.Graph)
	dist := make([][]int, n)

	// Initialization of the distance matrix
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			if g.Graph[i][j] == 0 && i != j {
				dist[i][j] = math.MaxInt32 // If there is no edge, we set infinity
			} else {
				dist[i][j] = g.Graph[i][j]
			}
		}
	}

	// Floyd-Warshall algorithm
	for k := range n {
		for i := range n {
			for j := range n {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}

	return dist
}
