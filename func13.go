package main

import (
	"errors"
	"strconv"
)

func sumDigits(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("O número não pode ser negativo")
	}

	strNumber := strconv.Itoa(number)
	sum := 0

	for _, digit := range strNumber {
		digitValue, _ := strconv.Atoi(string(digit))
		sum += digitValue
	}

	return sum, nil
}
