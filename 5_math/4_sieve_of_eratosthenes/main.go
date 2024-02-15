package main

import "fmt"

func main() {
	primes := sieve(100)
	fmt.Println(primes)
}

func sieve(n int) []int {
	mark := make([]bool, n)
	primes := []int{2}

	for i := 3; i*i <= n; i += 2 {
		if mark[i] == false {
			for j := i * i; j < n; j += i {
				mark[j] = true
			}
		}
	}
	for i := 3; i < n; i += 2 {
		if mark[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
