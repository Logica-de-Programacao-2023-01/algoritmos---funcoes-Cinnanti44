package main

import (
	"errors"
	"fmt"
)

func calcularMedia(lista []float64) (float64, error) {
	if len(lista) == 0 {
		return 0, errors.New("Erro: lista vazia")
	}

	soma := 0.0
	for _, num := range lista {
		soma += num
	}

	media := soma / float64(len(lista))
	return media, nil
}

func main() {
	numeros := []float64{2.5, 4.8, 6.1, 3.9}
	media, err := calcularMedia(numeros)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Média:", media)
	}

}
