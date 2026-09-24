package Conversor

import "fmt"

func ConvertirMoneda() {
	fmt.Println("=== CONVERSOR DE MONEDAS ===")
	var dolares float64
	var opcion int

	fmt.Print("Ingresa la cantidad en dólares ($): ")
	fmt.Scan(&dolares)

	fmt.Println("Elige la moneda (1: Euros, 2: Libras, 3: Won, 4: BTC): ")
	fmt.Scan(&opcion)

	if opcion == 1 {
		fmt.Printf("Equivale a %.2f Euros\n", dolares*0.92)
	} else if opcion == 2 {
		fmt.Printf("Equivale a %.2f Libras Esterlinas\n", dolares*0.79)
	} else if opcion == 3 {
		fmt.Printf("Equivale a %.2f Wones\n", dolares*1330.0)
	} else if opcion == 4 {
		fmt.Printf("Equivale a %.6f BTC\n", dolares*0.000016)
	} else {
		fmt.Println("Opción no válida")
	}
}
