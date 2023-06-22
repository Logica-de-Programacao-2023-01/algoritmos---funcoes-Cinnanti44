package main

import (
	"errors"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("O número é negativo")
	}

	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}
