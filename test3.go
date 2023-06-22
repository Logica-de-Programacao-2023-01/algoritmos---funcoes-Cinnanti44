package main

import (
	"errors"
	"fmt"
)

func contarCaracteres(str string) (int, error) {
	if len(str) == 0 {
		return 0, errors.New("Erro: string vazia")
	}

	return len(str), nil
}

func main() {
	texto := "Olá, mundo!"
	numCaracteres, err := contarCaracteres(texto)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Número de caracteres:", numCaracteres)
	}

	textoVazio := ""
	numCaracteres, err = contarCaracteres(textoVazio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Número de caracteres:", numCaracteres)
	}
}
