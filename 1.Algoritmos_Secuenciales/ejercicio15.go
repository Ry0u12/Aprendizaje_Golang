//Escribir un programa para calcular el área de un círculo:

package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	var a float64
	fmt.Print("Ingrese el radio del circulo: ")
	fmt.Scan(&r)
	a = 3.1416 * math.Pow(r, 2)
	fmt.Print("\n El area del circulo es: ", a)
}
