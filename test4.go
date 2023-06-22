package main

import (
	"errors"
	"fmt"
)

func encontrarMenorValor(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Erro: slice vazio")
	}

	menor := slice[0]
	for _, num := range slice {
		if num < menor {
			menor = num
		}
	}

	return menor, nil
}

func main() {
	valores := []int{8, 2, 5, 1, 9}
	menorValor, err := encontrarMenorValor(valores)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Menor valor:", menorValor)
	}

	valoresVazios := []int{}
	menorValor, err = encontrarMenorValor(valoresVazios)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Menor valor:", menorValor)
	}
}
