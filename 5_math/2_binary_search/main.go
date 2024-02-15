package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binary_search(1, arr))
	fmt.Println(binary_search(2, arr))
	fmt.Println(binary_search(3, arr))
	fmt.Println(binary_search(4, arr))
	fmt.Println(binary_search(5, arr))
	fmt.Println(binary_search(6, arr))
	fmt.Println(binary_search(0, arr))
}

// a
// m = l + (r-l)/2 = (l+r)/2
// a > m
// l = m+1
// a < m
// r = m-1
func binary_search(a int, arr []int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2
		if arr[m] == a {
			return m
		}
		if arr[m] < a {
			l = m + 1
			continue
		}
		if arr[m] > a {
			r = m - 1
			continue
		}
	}

	return -1
}
