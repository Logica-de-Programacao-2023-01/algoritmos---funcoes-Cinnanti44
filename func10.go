package main

import (
	"errors"
	"sort"
)

func OrdenarSlice(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("slice vazio")
	}

	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)

	return sorted, nil
}

func main() {
	// Exemplo de uso
	slice := []int{5, 2, 8, 1, 9, 3}
	sortedSlice, err := OrdenarSlice(slice)

	if err != nil {
		// Lidar com o erro
	} else {
		// Usar o slice ordenado
	}
}
