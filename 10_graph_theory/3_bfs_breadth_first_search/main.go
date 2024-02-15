package main

import (
	"fmt"
)

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
	bfsList(adjacencyList)
	fmt.Println()
	bfsMatrix(adjacencyMatrix)
}

func bfsList(graph [][]int) {
	startVertex := 0
	used := make([]int, len(graph))
	used[startVertex] = 1

	q := &IntHeap{}

	q.Push(startVertex)

	for q.Len() != 0 {
		currentVertex := q.Pop()
		fmt.Printf("%d ", currentVertex+1)

		for i := 0; i < len(graph[currentVertex]); i++ {
			adjacentVertex := graph[currentVertex][i]
			if used[adjacentVertex] == 0 {
				q.Push(adjacentVertex)
				used[adjacentVertex] = 1
			}
		}
	}
}
func bfsMatrix(graph [][]int) {

	startVertex := 0
	used := make([]int, len(graph))
	used[startVertex] = 1

	q := &IntHeap{}

	q.Push(startVertex)
	for q.Len() != 0 {
		currentVertex := q.Pop()
		fmt.Printf("%d ", currentVertex+1)
		for i := 0; i < len(graph[currentVertex]); i++ {
			if graph[currentVertex][i] == 1 {

				adjacentVertex := i
				if used[adjacentVertex] == 0 {
					q.Push(adjacentVertex)
					used[adjacentVertex] = 1
				}
			}
		}
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x int) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x)
}

func (h *IntHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[0]
	if n == 1 {
		*h = []int{}
	} else {
		*h = old[1:]
	}
	return x
}
