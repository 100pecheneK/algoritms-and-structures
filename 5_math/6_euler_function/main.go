package main

import "fmt"

// a, b - взаимопростые если gcd(a,b=1)
// f(n) = ... [1;n]
// f(9) = 6
// f(n)
// n = 84 = 2*2*3*7 -> 2, 3, 7 O(sqrt(n))
// 84 * (1-1/2) * (1-1/3) * (1-1/7) = 24
func main() {
	fmt.Println(euler(9), 6)
	fmt.Println(euler_fast(9), 6)
	fmt.Println(euler(36), 12)
	fmt.Println(euler_fast(36), 12)
	fmt.Println(euler(84), 24)
	fmt.Println(euler_fast(84), 24)
}

func euler(n int) int {
	count := 1
	for i := 2; i < n; i++ {
		if v := gcd(i, n); v == 1 {
			count++
		}
	}
	return count
}
func euler_fast(n int) int {
	result := n
	prime := 2
	for n >= prime*prime {
		if n%prime == 0 {
			result = result / prime * (prime - 1) // (1-1/p) = (p-1)/p
			for n%prime == 0 {
				n /= prime
			}
		}
		prime++
	}
	if n != 1 {
		result = result / n * (n - 1) // (1-1/p) = (p-1)/p
	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
