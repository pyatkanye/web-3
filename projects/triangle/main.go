package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Введите длину первого катета (a): ")
	fmt.Scan(&a)
	fmt.Print("Введите длину второго катета (b): ")
	fmt.Scan(&b)

	hypotenuse := math.Sqrt(a*a + b*b)

	fmt.Printf("Длина гипотенузы: %.2f\n", hypotenuse)
}
