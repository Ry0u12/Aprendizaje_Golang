/*Escribir un programa que declare una variable entera N y asígnale un valor. A
continuación, escribe las instrucciones que realicen los siguientes:
Incrementar N en 77. Decrementarla en 3. Duplicar su valor.*/

package main

import "fmt"

func main() {
	var N int64
	fmt.Scan(&N)
	N = N + 77
	fmt.Println(N)
	N = N - 3
	fmt.Println(N)
	N = N * 2
	fmt.Println(N)

}
