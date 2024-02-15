package main

import "fmt"

func main() {
	graph := [][]int{
		{1, 2},
		{0, 3, 6},
		{0, 3},
		{1, 2, 4, 5, 6},
		{3, 7},
		{3, 7},
		{1, 3, 7},
		{4, 5, 6, 8},
		{7},
	}

	bfsList(graph, 1)
}
func makeVector(l int) []int {
	v := make([]int, 0, l)
	for i := 0; i < l; i++ {
		v = append(v, -1)
	}
	return v
}
func bfsList(graph [][]int, startVertex int) {
	used := makeVector(len(graph))

	used[startVertex] = 0

	q := &IntHeap{}

	q.Push(startVertex)

	for q.Len() != 0 {
		currentVertex := q.Pop()
		fmt.Printf("%d ", currentVertex+1)

		for i := 0; i < len(graph[currentVertex]); i++ {
			adjacentVertex := graph[currentVertex][i]
			if used[adjacentVertex] == -1 {
				q.Push(adjacentVertex)
				used[adjacentVertex] = used[currentVertex] + 1
			}
		}
	}
	fmt.Println()
	for i, v := range used {
		fmt.Println(i+1, v)
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
