package main

import "fmt"

func main() {
	fmt.Println(51 & 25) // 17
	fmt.Println(51 | 25) // 59
	fmt.Println(51 ^ 25) // 42
	fmt.Println(^uint(23))
	fmt.Println(12 << 2) // 48
	fmt.Println(12 >> 2) // 3
	fmt.Println(1 << 4)  // 16
	fmt.Println(1 << 5)  // ...
	fmt.Println(1 << 6)  // ...
	fmt.Println(1 << 7)  // ...
	fmt.Println(1 << 8)  // 256

	n := 17           // 10001
	i := 2            //   1
	n = n | (1 << i)  // 10101
	fmt.Println(n)    //   0
	n = n & ^(1 << i) // 10001
	fmt.Println(n)
	n = n ^ (1 << i) // 1 <- 0 | 0 <- 1
	fmt.Println(n)

	if n&1<<i != 0 {
		fmt.Printf("%d bit is not zero %8b", i, n)
	}

	fmt.Println()
	singleNumber([]int{3, 5, 2, 2, 3})
	// 0000 ^ 0011 = 0011
	// 0101 ^ 0011 = 0110
	// 0010 ^ 0110 = 0100
	// 0010 ^ 0100 = 0110
	// 0011 ^ 0110 = 0101 (5)

	singleNumber([]int{2, 3, 4, 1, 1, 3, 4})
}

func singleNumber(arr []int) int {
	res := 0
	for _, v := range arr {
		res ^= v
	}
	fmt.Println("res:", res)
	return res
}
