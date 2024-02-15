package main

import "fmt"

func main() {
	fmt.Println(gcd(15, 28))

	fmt.Println(gcd(16, 24))
	fmt.Println(gcd(24, 16))
	fmt.Println(gcd(64, 128))
	fmt.Println(lcm(15, 28))
	fmt.Println(lcm(16, 24))
	fmt.Println(lcm(24, 16))
	fmt.Println(lcm(64, 128))
}

// наибольший общий делитель
// f(a,b) = {
// if b == 0, a
// if f(b, a % b)
// }
// O(log(min(a,b)))
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// наименьшее общее кратное
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
