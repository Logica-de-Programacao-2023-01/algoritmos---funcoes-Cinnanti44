package main

import "fmt"

type triangulo struct {
	altura float64
	base   float64
}

func calculararea(p triangulo) float64 {
	return p.altura * p.base / 2

}
func main() {
	p := triangulo{
		base:   5.0,
		altura: 3.0,
	}
	area := calculararea(p)
	fmt.Println("Área do triangulo:", area)

}
