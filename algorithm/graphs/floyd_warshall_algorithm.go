package graphs

import "math"

type GraphFloydWarshall struct {
	Graph [][]int
}

func FloydWarshall(g GraphFloydWarshall) [][]int {
	n := len(g.Graph)
	dist := make([][]int, n)

	// Инициализация матрицы расстояний
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if g.Graph[i][j] == 0 && i != j {
				dist[i][j] = math.MaxInt32 // Если нет ребра, ставим бесконечность
			} else {
				dist[i][j] = g.Graph[i][j]
			}
		}
	}

	// Алгоритм Флойда-Уоршелла
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}

	return dist
}
