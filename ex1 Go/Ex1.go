package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite os coeficientes a, b e c: ")
	fmt.Scan(&a, &b, &c)

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("A equação nao possui raizes reais.")
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("As raizes da equacao sao: x1 = %.2f e x2 = %.2f\n", x1, x2)
	}
}