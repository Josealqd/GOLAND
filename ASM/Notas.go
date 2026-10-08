package main

import "fmt"

func main() {
	// 1. Array bidimensional de 6 estudiantes y 4 notas
	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 8.0},
		{6.0, 7.5, 8.0, 9.0},
		{10.0, 9.5, 9.0, 9.8},
		{5.5, 6.0, 7.0, 6.5},
		{8.0, 8.0, 8.0, 8.0},
		{7.5, 8.5, 9.0, 7.0},
	}

	var sumaClase float64
	var totalNotas int

	fmt.Println("=== NOTAS DE LOS ESTUDIANTES ===")

	for i := 0; i < 6; i++ {
		
		sliceEstudiante := notas[i][:]

		sumaEstudiante := 0.0
		max := sliceEstudiante[0]
		min := sliceEstudiante[0]

		for j := 0; j < len(sliceEstudiante); j++ {
			nota := sliceEstudiante[j]
			
			sumaEstudiante = sumaEstudiante + nota

			if nota > max {
				max = nota
			}
			
			if nota < min {
				min = nota
			}


			sumaClase = sumaClase + nota
			totalNotas++
		}

		promedio := sumaEstudiante / 4.0 

		fmt.Printf("Estudiante %d -> Promedio: %.2f | Max: %.2f | Min: %.2f\n", i+1, promedio, max, min)
	}

	promedioGeneral := sumaClase / float64(totalNotas)
	fmt.Printf("\nPromedio de toda la clase: %.2f\n", promedioGeneral)
}
