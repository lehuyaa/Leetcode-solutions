package main

func countPrimes(n int) int {
	if n == 0 {
		return 0
	}
	isPrime := make([]bool, n+1)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i < n; i++ {
		if isPrime[i] {
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 0; i <= n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
