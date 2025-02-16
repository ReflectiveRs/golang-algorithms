package main

import (
	algorithmsGraphs "algorithms/algorithm/graphs"
	algorithmsSearch "algorithms/algorithm/search"
	algorithmsSorting "algorithms/algorithm/sorting"
	algorithmsTree "algorithms/algorithm/tree"
	"fmt"
	"log"
	"math"
)

func main() {
	// example sorting
	var arrForSort = []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("QuickSort Исходный массив:", arrForSort)
	var sortedArr = algorithmsSorting.QuickSort(arrForSort)
	fmt.Println("QuickSort Отсортированный массив:", sortedArr)

	// example search
	var arrForSearch = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var targetForSearch = 7
	var indexSearchElement = algorithmsSearch.BinarySearch(arrForSearch, targetForSearch)

	if indexSearchElement != -1 {
		fmt.Printf("Результат BinarySearch Элемент %d найден на индексе %d\n", targetForSearch, indexSearchElement)
	} else {
		fmt.Printf("Результат BinarySearch Элемент %d не найден\n", targetForSearch)
	}

	// example dijkstra
	var graphDijkstra = &algorithmsGraphs.GraphDijkstra{
		Vertices: 5,
		Edges: map[int][]algorithmsGraphs.EdgeDijkstra{
			0: {{To: 1, Weight: 10}, {To: 2, Weight: 3}},
			1: {{To: 2, Weight: 1}, {To: 3, Weight: 2}},
			2: {{To: 1, Weight: 4}, {To: 3, Weight: 8}, {To: 4, Weight: 2}},
			3: {{To: 4, Weight: 7}},
			4: {},
		},
	}

	var startVertexForDijkstra = 0
	distancesGraphDijkstra, err := graphDijkstra.Dijkstra(startVertexForDijkstra)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Dijkstra Кратчайшие расстояния от вершины %d:\n", startVertexForDijkstra)
	fmt.Printf("Dijkstra ")
	for i, distance := range distancesGraphDijkstra {
		if distance == math.MaxInt32 {
			fmt.Printf("Вершина %d: недоступна, ", i)
		} else {
			fmt.Printf("Вершина %d: %d, ", i, distance)
		}
	}
	fmt.Println()

	// example floydWarshall
	var graphFloydWarshall = algorithmsGraphs.GraphFloydWarshall{
		Graph: [][]int{
			{0, 3, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0},
			{0, 0, 0, 7, 0, 2},
			{0, 0, 0, 0, 2, 0},
			{6, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
	}

	var graphFloydWarshallDistances = algorithmsGraphs.FloydWarshall(graphFloydWarshall)
	fmt.Println("FloydWarshall Матрица кратчайших расстояний:")
	for i := 0; i < len(graphFloydWarshallDistances); i++ {
		for j := 0; j < len(graphFloydWarshallDistances[i]); j++ {
			if graphFloydWarshallDistances[i][j] == math.MaxInt32 {
				fmt.Print("Infinity ")
			} else {
				fmt.Print(graphFloydWarshallDistances[i][j], " ")
			}
		}
		fmt.Println()
	}
	fmt.Println("End FloydWarshall")

	// example depth first search
	graphDFS := algorithmsGraphs.NewGraphDFS()
	graphDFS.AddEdge(0, 1)
	graphDFS.AddEdge(0, 2)
	graphDFS.AddEdge(1, 3)
	graphDFS.AddEdge(1, 4)
	graphDFS.AddEdge(2, 5)
	graphDFS.AddEdge(2, 6)

	// Recursive DFS
	visitedGraphDFS := make(map[int]bool)
	var resultGraphDFSRecursive []int
	graphDFS.DFSRecursive(0, visitedGraphDFS, &resultGraphDFSRecursive)
	fmt.Println("Результат DFSRecursive:", resultGraphDFSRecursive)

	// Iterative DFS
	resultGraphDFSIterative := graphDFS.DFSIterative(0)
	fmt.Println("Результат DFSIterative:", resultGraphDFSIterative)

	// example breadth_first_search
	graph := algorithmsGraphs.NewGraphBFS()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	resultGraphBFS := graph.BFS(0)
	fmt.Println("Результат BFS:", resultGraphBFS)

	// example binary_tree
	bst := &algorithmsTree.BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(12)
	bst.Insert(18)

	fmt.Println("binary_tree Поиск 7:", bst.Search(7))
	fmt.Println("binary_tree Поиск 20:", bst.Search(20))
	fmt.Println("binary_tree Симметричный обход:", bst.InOrderTraversal())

	// example AVL tree
	treeAVL := &algorithmsTree.AVLTree{}
	valuesAVL := []int{30, 20, 40, 10, 25, 35, 50}

	for _, v := range valuesAVL {
		treeAVL.Insert(v)
	}

	fmt.Println("AVLTree Симметричный обход InOrder:", treeAVL.InOrderTraversal())
	fmt.Println("AVLTree Поиск 25:", treeAVL.Search(25))
	fmt.Println("AVLTree Поиск 100:", treeAVL.Search(100))

	// example red black tree
	treeRBT := &algorithmsTree.RedBlackTree{}

	treeRBT.Insert(10)
	treeRBT.Insert(20)
	treeRBT.Insert(30)
	treeRBT.Insert(15)
	treeRBT.Insert(25)
	treeRBT.Insert(5)

	// Обход дерева in-order
	resultRBT := treeRBT.InOrderTraversal()
	fmt.Println("RedBlackTree In-order traversal:", resultRBT)

	// Поиск элемента
	fmt.Println("RedBlackTree Поиск 15:", treeRBT.Search(15))
	fmt.Println("RedBlackTree Поиск 40:", treeRBT.Search(40))
}
