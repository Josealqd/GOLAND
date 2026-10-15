package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese una frase o texto: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	texto = strings.ToLower(texto)
	palabras := strings.Fields(texto)

	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	var palabraMasRepetida string
	maxRepeticiones := 0

	fmt.Println("\n--- Resultados del conteo ---")
	
	for palabra, cantidad := range conteo {
		fmt.Printf("La palabra '%s' aparece %d vez/veces.\n", palabra, cantidad)

		if cantidad > maxRepeticiones {
			maxRepeticiones = cantidad
			palabraMasRepetida = palabra
		}
	}

	fmt.Printf("\nLa palabra que más se repite es '%s' con %d repeticiones.\n", palabraMasRepetida, maxRepeticiones)
}