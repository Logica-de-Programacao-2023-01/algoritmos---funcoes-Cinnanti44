package main

import "fmt"

type Livro struct {
	titulo          string
	autor           string
	anoDePublicacao int
}

func imprimirLivro(p Livro) {
	fmt.Println("Título:", p.titulo)
	fmt.Println("Autor:", p.autor)
	fmt.Println("Ano de Publicação:", p.anoDePublicacao)
}

func main() {
	p := Livro{
		titulo:          "A Guerra dos Tronos",
		autor:           "George R. R. Martin",
		anoDePublicacao: 1996,
	}

	imprimirLivro(p)
}
