package main

import "fmt"

func main() {
	var x int = 42
	var ptr *int = &x
	fmt.Println("Valor de x:",
		*ptr)

}
