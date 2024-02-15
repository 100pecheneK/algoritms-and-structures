package main

import (
	"fmt"
)

func search(a []int, x int) int { // O(n)
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		}
	}
	return -1
}

func swap(a, b *int) { // O(1)
	a, b = b, a
}

func cycle_l() { // O(1)
	c := 4

	for i := 0; i < c; i++ {
	}
}

func cycle_n(n int) { // O(n)
	for i := 0; i < n; i++ {
	}
}

func cycle_nm1(n, m int) { // O(max(n,m) = O(n+m)
	for i := 0; i < n; i++ {
	}
	for i := 0; i < m; i++ {
	}
}
func cycle_nm2(n, m int) { // O(n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
		}
	}
}

func cycle_nm3(n int) { // O(n^2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
		}
	}
}
func cycle_nm4(n int) { // O(n^3)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
			}
		}
	}
}

func cycle_ij(n int) { // O(n^2)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
		}
	}
	// n + (n-1) + (n-2) + ... + 1 == n * (n+1)/2
	// n*(n+1)/2 = 1/2 * (n*n + n) = O(n^2)
}

func cycle_n123(n int) { // O(n^3)
	for i := 0; i < n; i++ { // O(n)
	}
	for i := 0; i < n; i++ { // O(n^2)
		for j := 0; j < n; j++ {
		}
	}

	for i := 0; i < n; i++ { // O(n^3)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
			}
		}
	}
}
func cycle_abc123(a, b, c int) { // O(a+b^2+c^3)
	for i := 0; i < a; i++ { // O(n)
	}
	for i := 0; i < b; i++ { // O(n^2)
		for j := 0; j < b; j++ {
		}
	}

	for i := 0; i < c; i++ { // O(n^3)
		for j := 0; j < c; j++ {
			for k := 0; k < c; k++ {
			}
		}
	}
}

func cycle_log1(a int) int { // O(log(a))
	sum := 0
	for a != 0 { // log10(a)
		sum += a % 10
		a /= 10
	}
	return sum
}

func cycle_log2(n int) int { // O(n * log(a))
	sum := 0
	for i := 0; i < n; i++ {
		a := i

		for a != 0 { // log10(a)
			sum += a % 10
			a /= 10
		}
	}
	return sum
}
func cycle_sqrt(n int) { // O(sqrt(a))
	for i := 0; i*i < n; i++ {
		fmt.Println(i)
	}
}

// O(2^n), O(3^n), O(n!)
// 10^9 - 1 сек в среднем
func main() {}
