package graphs

import (
	"container/heap"
	"errors"
	"math"
)

var (
	ErrWeightIsNegative = errors.New("dijkstra algorithms, the weight cannot be a negative value")
)

// Graph is represented as a structure containing an adjacency list
type GraphDijkstra struct {
	Vertices int
	Edges    map[int][]EdgeDijkstra
}

// Edge represents a structure for the edges of a graph
type EdgeDijkstra struct {
	To     int
	Weight int
}

// PriorityQueue implements a priority queue data structure
type PriorityQueue []*ItemDijkstra

// Item represents an element in a priority queue
type ItemDijkstra struct {
	vertex   int
	distance int
	index    int // Index in queue
}

// Len, Less, Swap, Push и Pop реализуют интерфейс heap.Interface
func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ItemDijkstra)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra performs Dijkstra's algorithm
func (g *GraphDijkstra) Dijkstra(start int) ([]int, error) {
	distances := make([]int, g.Vertices)
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &ItemDijkstra{vertex: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*ItemDijkstra)

		for _, edge := range g.Edges[current.vertex] {
			newDist := distances[current.vertex] + edge.Weight
			if newDist < 0 {
				return nil, ErrWeightIsNegative
			}
			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				heap.Push(pq, &ItemDijkstra{vertex: edge.To, distance: newDist})
			}
		}
	}

	return distances, nil
}
