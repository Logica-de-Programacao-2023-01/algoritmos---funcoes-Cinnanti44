package main

import (
	"errors"
	"math"
)

func VerificarPrimo(number int) (bool, error) {
	if number < 0 {
		return false, errors.New("número negativo")
	}

	if number < 2 {
		return false, nil
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	// Exemplo de uso
	num := 13
	isPrime, err := VerificarPrimo(num)

	if err != nil {
		// Lidar com o erro
	} else {
		if isPrime {
			// O número é primo
		} else {
			// O número não é primo
		}
	}
}
