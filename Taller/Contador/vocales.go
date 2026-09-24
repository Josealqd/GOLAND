package Contador

import "fmt"

func ContarVocales() {
	fmt.Println("=== CONTADOR DE VOCALES ===")
	var frase string
	fmt.Print("Ingresa una palabra o texto (sin espacios): ")
	fmt.Scan(&frase)

	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, letra := range frase {
		if letra == 'a' || letra == 'A' {
			a++
		} else if letra == 'e' || letra == 'E' {
			e++
		} else if letra == 'i' || letra == 'I' {
			i++
		} else if letra == 'o' || letra == 'O' {
			o++
		} else if letra == 'u' || letra == 'U' {
			u++
		}
	}

	fmt.Printf("A: %d\n", a)
	fmt.Printf("E: %d\n", e)
	fmt.Printf("I: %d\n", i)
	fmt.Printf("O: %d\n", o)												
	fmt.Printf("U: %d\n", u)
}
