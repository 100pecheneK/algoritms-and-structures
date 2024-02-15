package main

import "fmt"

func main() {
	fib_m := [][]int64{
		{1, 1},
		{1, 0},
	}
	fib_1 := [][]int64{
		{1, 0}, // F1, F0
	}

	var f1, f0 int64 = 1, 0
	var mod int64 = 10000

	for i := 0; i < 10; i++ {
		// fmt.Printf("F(%d) = %d\n", i, fib_1[0][0])
		fib_1 = matrix_mul(fib_1, fib_m, mod)

		var f2 int64 = (f1 + f0) % mod
		f0 = f1
		f1 = f2
		if f1 != fib_1[0][0] {
			fmt.Println("EEERROR:", f1, fib_1[0][0])
		}
		var matrix_pow [][]int64 = bin_pow(fib_m, i, mod)
		fibs := matrix_mul([][]int64{{1, 1}}, matrix_pow, mod)

		if fibs[0][0] != fib_1[0][0] {
			fmt.Println("ERROR:", fibs[0][0], fib_1[0][0])
		}
	}
	var matrix_pow [][]int64 = bin_pow(fib_m, 123456789123456789, mod)
	fibs := matrix_mul([][]int64{{1, 1}}, matrix_pow, mod)

	fmt.Println(fibs[0][0])
}

func createMatrix(n, m int) [][]int64 {
	rows := make([][]int64, n)
	for i := 0; i < n; i++ {
		rows[i] = make([]int64, m)
	}
	return rows
}
func matrix_mul(a, b [][]int64, mod int64) [][]int64 {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	c := createMatrix(n, k)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += (a[i][k] * b[k][j]) % mod
			}
			c[i][j] %= mod
		}
	}
	return c
}

func bin_pow(a [][]int64, n int, mod int64) [][]int64 {
	result := createMatrix(len(a), len(a))
	for i := 0; i < len(a); i++ {
		result[i][i] = 1
	}
	for n != 0 {
		if n&1 == 1 {
			result = matrix_mul(result, a, mod)
		}
		n >>= 1 // n/2
		a = matrix_mul(a, a, mod)
	}
	return result
}
