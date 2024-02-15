package main

import "fmt"

func main() {
	graph := [][]int{
		{4, 5},
		{3},
		{3},
		{1, 2, 9},
		{0},
		{0},
		{},
		{8},
		{7},
		{3},
	}

	used := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if used[i] == 0 {
			dfsList(graph, used, i)
			fmt.Println()
		}

	}
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
