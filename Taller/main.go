package main

import (
	"Taller/Contador"
	"Taller/Conversor"
	"fmt"
)

func main() {
	var opcion int

	fmt.Println("=== MENÚ PRINCIPAL ===")
	fmt.Println("1. Contador de Vocales")
	fmt.Println("2. Conversor de Monedas")
	fmt.Print("Elige una opción: ")
	fmt.Scan(&opcion)

	if opcion == 1 {
		Contador.ContarVocales()
	} else if opcion == 2 {
		Conversor.ConvertirMoneda()
	} else {
		fmt.Println("Opción incorrecta")
	}
}
