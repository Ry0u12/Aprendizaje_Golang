// Escribir un programa que cálcule del promedio de tres números
package main

import "fmt"

func main() {
	var num1 int64
	var num2 int64
	var num3 int64
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)
	fmt.Print("Ingrese el tercer numero: ")
	fmt.Scanln(&num3)
	promedio := (num1 + num2 + num3) / 3
	fmt.Println("El promedio de estos numeros es: ", promedio)
}
