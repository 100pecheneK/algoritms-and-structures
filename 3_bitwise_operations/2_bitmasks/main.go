package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	n := len(arr)

	for mask := 0; mask < 1<<n; mask++ {
		first := true
		fmt.Print("{")

		for i, v := range arr {
			if (mask & (1 << i)) != 0 {
				if !first {
					fmt.Print(", ")
				}
				first = false
				fmt.Print(v)
				continue
			}
		}
		fmt.Println("}")
	}
}
