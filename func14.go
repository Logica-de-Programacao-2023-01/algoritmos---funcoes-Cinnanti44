package main

import (
	"errors"
)

func intersection(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	numbers := make([]int, 0)
	numbersMap := make(map[int]bool)

	for _, num := range slice1 {
		numbersMap[num] = true
	}

	for _, num := range slice2 {
		if numbersMap[num] {
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}
