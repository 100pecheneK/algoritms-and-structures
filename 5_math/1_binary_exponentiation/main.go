package main

import "fmt"

func main() {
	fmt.Println(bin_pow(2, 10) == 1024)
	fmt.Println(bin_pow(3, 5) == 243)
	fmt.Println(bin_pow(3, 15) == 14348907)

	fmt.Println(bin_pow_fast(2, 10) == 1024)
	fmt.Println(bin_pow_fast(3, 5) == 243)
	fmt.Println(bin_pow_fast(3, 15) == 14348907)
}

// exp(a, n) = {
// if n odd, exp(a, n-1)*a
// if n even, exp(a, n/2)^2
// in n == 0, 1
// }
// O(n*log(n))
func bin_pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return bin_pow(a, n-1) * a
	}
	b := bin_pow(a, n/2)
	return b * b
}

func bin_pow_fast(a, n int) int {
	result := 1
	for n != 0 {
		// last bit == 1 => odd
		if n&1 == 1 {
			result *= a
		}
		n >>= 1 // n/2
		a *= a  // a, a^2, a^4, a^8
	}
	return result
}
