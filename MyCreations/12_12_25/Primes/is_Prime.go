package main

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n != 1
}

func Primes(n int) []int {
	result := []int{}
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}
	return result
}
