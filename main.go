package main

import (
	"fmt"
	"log"
)

func main() {
	num := 9
	var suma = 0

	if num > 0 {
		log.Printf("Var: es mayor a 0")
	} else if num < 0 {
		log.Printf("Var: es número negativo")
	} else {
		log.Printf("Var: es número positivo")
	}

	for i := 0; i < 10; i++ {
		suma += i
		fmt.Println(suma)
	}
}
