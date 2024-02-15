package main

import "fmt"

func main() {
	arr := []int{1, 7, 5, 3, 2, 3, 8, 9}
	k := 4
	sum := 0
	for i := 0; i < k; i++ { // O(k)
		sum += arr[i]
	}

	fmt.Println(sum)
	for i := 1; i <= len(arr)-k; i++ { // O(n)
		sum = sum - arr[i-1] + arr[i+k-1]
		fmt.Println(sum)
	}

	fmt.Println("---")
	for i := 0; i <= len(arr)-k; i++ { // O(n*k)
		sum := 0
		for j := 0; j < k; j++ {
			sum += arr[i+j]
		}
		fmt.Println(sum)

	}
}
