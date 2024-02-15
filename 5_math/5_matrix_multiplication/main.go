package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{3, 1, 2},
	}
	b := [][]int{
		{1, 2},
		{3, 2},
		{1, 2},
	}
	c := [][]int{
		{10, 12},
		{8, 12},
	}

	fmt.Println(matrix_mul(a, b), c)
}

func createMatrix(n, m int) [][]int {
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		rows[i] = make([]int, m)
	}
	return rows
}
func matrix_mul(a, b [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	c := createMatrix(n, k)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
