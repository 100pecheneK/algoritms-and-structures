package main

import "fmt"

// O(n*k)
func main() {
	n, k := 10, 2
	l := []int{1}

	for i := 1; i < n; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			if len(l)-1-j < 0 {
				break
			}
			sum += l[len(l)-1-j]
			sum %= 10000
		}
		l = append(l, sum)
	}
	fmt.Println(l)
}
