package main

import "fmt"

func divide(x, y float64) (float64, error) {

	if y == 0 {

		return 0, fmt.Errorf("cannot divide by zero")

	}
	return x / y, nil

}

func main() {
	result, err := divide(10, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
