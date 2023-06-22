package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func calcularmedia(p Aluno) (notas float64) {
	soma := 0.0
	for _, num := range p.notas {
		soma += num
	}

	media := soma / float64(len(p.notas))
	return media
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{8.5, 7.2, 9.0, 6.8},
	}

	media := calcularmedia(aluno)
	fmt.Println("Média das notas:", media)
}
