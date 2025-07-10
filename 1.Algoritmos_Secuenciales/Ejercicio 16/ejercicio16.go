// Escribir un programa que convierta grados Celsius a Fahrenheit:
package main

import "fmt"

func main() {
	var celcius float64
	fmt.Println("Ingresa el valor de los grados celsius: ")
	fmt.Scan(&celcius)
	var fahrenheit = celcius*(9.0/5.0) + 32.0
	fmt.Println("El valor en la escala farenheit es: ", fahrenheit)
}
