package main

import (
	"fmt"
	"io"
	"strings"
)

type Edges struct {
	a, b int
}

func main() {
	fmt.Println("Input:\nn=5 m=7\n1 2\n1 3\n1 4\n1 5\n2 3\n3 4\n4 5")
	in := strings.NewReader(`
	5 7
	1 2
	1 3
	1 4
	1 5
	2 3
	3 4
	4 5
	5 7
	1 2
	1 3
	1 4
	1 5
	2 3
	3 4
	4 5
	5 7
	1 2
	1 3
	1 4
	1 5
	2 3
	3 4
	4 5
	`)

	fmt.Println("List of edges")
	listOfEdges(in)
	fmt.Println("Adjacency matrix")
	adjacencyMatrix(in)
	fmt.Println("Adjacency list")
	adjacencyList(in)
}

func listOfEdges(in io.Reader) {
	// List of edges
	var n, m int
	fmt.Fscan(in, &n, &m)
	edges := make([]*Edges, m)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		edges[i] = &Edges{a, b}
	}

	for i := 0; i < m; i++ {
		fmt.Println(edges[i].a, edges[i].b)
	}

}

func adjacencyMatrix(in io.Reader) {
	// Adjacency matrix
	var n, m int
	fmt.Fscan(in, &n, &m)
	matrix := _matrix(n, n)

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		// 1 <= a,b <= n
		a--
		b--
		matrix[a][b] = 1
		// если b не сосед a, то следующая строка не нужна
		matrix[b][a] = 1
	}

	fmt.Print("   ")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, " ")
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(i+1, "| ")
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}

func adjacencyList(in io.Reader) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	matrix := make([][]int, n)

	var a, b int

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		a--
		b--
		matrix[a] = append(matrix[a], b)
		// если b не сосед a, то следующая строка не нужна
		matrix[b] = append(matrix[b], a)
	}

	for i, r := range matrix {
		fmt.Print(i+1, "| ")
		for _, c := range r {
			fmt.Print(c+1, " ")
		}
		fmt.Println()
	}
}

func _matrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	return matrix
}
