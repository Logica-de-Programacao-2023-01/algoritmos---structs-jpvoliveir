package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func calcularArea(t Triangulo) float64 {
	area := t.Base * t.Altura / 2
	return area
}

func main() {
	triangulo := Triangulo{Base: 5.0, Altura: 3.0}
	area := calcularArea(triangulo)
	fmt.Printf("A área do triângulo é: %.2f\n", area)
}
