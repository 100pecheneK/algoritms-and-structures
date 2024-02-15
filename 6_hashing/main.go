package main

import "fmt"

func main() {
	s1 := "aaa"
	s2 := "abb"
	s3 := "abb"

	fmt.Println(string_hash(s1))
	fmt.Println(string_hash(s2))
	fmt.Println(string_hash(s3))
	fmt.Println(string_hash(s3) == string_hash(s2))

	s := "qwertyuiopasdfghjklzxcvbnm"
	hashes := make([]int64, len(s)+1)
	hashes[0] = 0

	primes := make([]int64, len(s)+1)
	primes[0] = 1

	var p int64 = 301
	var m int64 = 1e9 + 7

	for i := 0; i < len(s); i++ {
		hashes[i+1] = hashes[i]*p + int64(s[i]-'a') + 1
		hashes[i+1] %= m
		primes[i+1] = primes[i] * p
		primes[i+1] %= m
	}

	i, j := 3, 7 //...rtyui...
	substr_hash := (hashes[j+1] - (hashes[i]*primes[j-i+1])%m) % m
	if substr_hash < 0 {
		substr_hash += m
	}
	fmt.Println(substr_hash) // 301429957

}

// O(n)
func string_hash(s string) int64 {
	// count small english letters
	var p int64 = 31
	var m int64 = 1e9 + 7
	var hash int64
	var prime_pow int64 = 1

	for _, r := range s {
		hash = (hash + (int64(r-'a')+1)*prime_pow) % m
		prime_pow = (prime_pow * p) % m
	}
	return hash
}
