package main

import (
	"fmt"
	"practica/saludo"
)

func main() {
	fmt.Println("***Bienvenidos a la clase de paquetes 😊😊😊***")
	mensaje:=saludo.Saludar("Juan")
	fmt.Println(mensaje)
}
