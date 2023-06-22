package main

import (
	"errors"
)

type IntOperation func(int) int

func applyAndSum(numbers []int, operation IntOperation) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	sum := 0

	for _, num := range numbers {
		result := operation(num)
		sum += result
	}

	return sum, nil
}
