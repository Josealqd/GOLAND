package main

import "fmt"

func sacarGanador(votos map[string]int) string {
	actividadGanadora := ""
	maxVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			actividadGanadora = actividad
		}
	}
	return actividadGanadora
}

func main() {
	
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("Opciones para votar: deportes, videojuegos, cine, musica")
	fmt.Println("Por favor, escribe exactamente la palabra.")


	for i := 1; i <= 5; i++ {
		var voto string
		fmt.Printf("Ingresa tu voto %d: ", i)
		fmt.Scan(&voto)

	
		_, existe := votos[voto]
		
		if existe {
			votos[voto] = votos[voto] + 1 
		} else {
			fmt.Println("Esa opción no existe. Perdiste tu voto.")
		}
	}

	
	fmt.Println("\n--- RESULTADOS ---")
	for actividad, cantidad := range votos {
		fmt.Printf("%s: %d votos\n", actividad, cantidad)
	}

	// Llamar a la función del paso 5
	ganador := sacarGanador(votos)
	fmt.Printf("\nLa actividad con más votos fue: %s\n", ganador)
}