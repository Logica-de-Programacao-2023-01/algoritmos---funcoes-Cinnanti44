package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func calcularArea(p Retangulo) float64 {
	return p.largura * p.altura
}

func main() {
	p := Retangulo{
		largura: 5.0,
		altura:  3.0,
	}

	area := calcularArea(p)
	fmt.Println("Área do retângulo:", area)
}
