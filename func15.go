package main

import (
	"errors"
)

func isNumberPresent(number int, numbers []int) (bool, error) {
	if len(numbers) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, num := range numbers {
		if num == number {
			return true, nil
		}
	}

	return false, nil
}
