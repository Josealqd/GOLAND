package main

import "fmt"

func averageGrade() {
	var n int
	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Cantidad de estudiantes no válida.")
		return
	}

	var suma float64 = 0

	for i := 1; i <= n; i++ {
		var nota float64
		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", i)
		fmt.Scan(&nota)
		suma += nota
	}

	promedio := suma / float64(n)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("Estado: Aprobado")
	} else {
		fmt.Println("Estado: Reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Mensaje: Excellent performance")
	case promedio >= 80 && promedio <= 89:
		fmt.Println("Mensaje: Good performance")
	case promedio >= 70 && promedio <= 79:
		fmt.Println("Mensaje: Satisfactory performance")
	default:
		fmt.Println("Mensaje: Needs improvement")
	}
}

func sumarHastaN() {
	var n int
	fmt.Print("Ingrese un número n: ")
	fmt.Scan(&n)

	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}

	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, suma)
}

func celsiusAFahrenheit() {
	var c float64
	fmt.Print("Ingrese la temperatura en grados Celsius: ")
	fmt.Scan(&c)

	f := (c * 9 / 5) + 32
	fmt.Printf("%.2f °C equivalen a %.2f °F\n", c, f)
}

func fahrenheitACelsius() {
	var f float64
	fmt.Print("Ingrese la temperatura en grados Fahrenheit: ")
	fmt.Scan(&f)

	c := (f - 32) * 5 / 9
	fmt.Printf("%.2f °F equivalen a %.2f °C\n", f, c)
}

func main() {
	opcion := ""

	for opcion != "0" && opcion != "salir" {
		fmt.Println("\n==============================")
		fmt.Println("        MENÚ PRINCIPAL        ")
		fmt.Println("==============================")
		fmt.Println("1. Calcular promedio de notas")
		fmt.Println("2. Sumar números del 1 al n")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0 o 'salir' para terminar")
		fmt.Print("Elija una opción: ")

		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Programa finalizado. Hasta luego")
		} else {
			switch opcion {
			case "1":
				averageGrade()
			case "2":
				sumarHastaN()
			case "3":
				celsiusAFahrenheit()
			case "4":
				fahrenheitACelsius()
			default:
				fmt.Println("Opción inválida. Por favor, intente de nuevo.")
			}
		}
	}
}
