package main

import "fmt"

func main() {
	adjacencyList := [][]int{
		{1, 3},
		{0, 5, 6},
		{6},
		{0, 4},
		{3, 7},
		{1, 6},
		{1, 2, 5},
		{4},
	}
	adjacencyMatrix := [][]int{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
	}
	n := 8
	used := make([]int, n)
	dfsList(adjacencyList, used, 0)
	fmt.Println()
	used = make([]int, n)
	dfsMatrix(adjacencyMatrix, used, 0)
}

func dfsList(graph [][]int, used []int, currrentVertex int) {
	fmt.Println(currrentVertex + 1)
	used[currrentVertex] = 1
	for i := 0; i < len(graph[currrentVertex]); i++ {
		adjacentVertex := graph[currrentVertex][i]
		if used[adjacentVertex] != 1 {
			dfsList(graph, used, adjacentVertex)
		}
	}
}
func dfsMatrix(graph [][]int, used []int, currentVertex int) {
	fmt.Println(currentVertex + 1)
	used[currentVertex] = 1
	for i := 0; i < len(graph[currentVertex]); i++ {
		if graph[currentVertex][i] == 1 {
			if used[i] != 1 {
				dfsMatrix(graph, used, i)
			}
		}
	}
}
