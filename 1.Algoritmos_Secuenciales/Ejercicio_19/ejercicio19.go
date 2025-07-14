/*
José y Pedro están caminando por el campus de la EMI, ellos colocaron su dinero en una sola
billetera, todo está mezclado, ahora necesitan saber cuánto dinero tenía José y cuánto Pedro.
Solo recuerdan que en total hay X monedas y la diferencia entre su dinero es Y, además
recuerdan que José tenía más dinero. Se pide realizar un algoritmo para resolverlo, aplicar los
pasos necesarios
*/
package main

import "fmt"

func main() {
	var x float64
	var y float64
	fmt.Print("Ingrese la cantidad de dinero que hay en total (x): ")
	fmt.Scanln(&x)
	fmt.Print("Ingrese la cantidad de dinero que hay de diferencia (y): ")
	fmt.Scanln(&y)
	Jose := (x - y) / 2.0
	Pedro := y + Jose
	fmt.Print("El dinero que tenia Pedro es: ", Pedro, ", y el dinero que tenia Jose es: ", Jose)
}
