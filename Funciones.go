package main

import "fmt"

/*

FUN <nombre>(param1,param2, ....param n)<valores de retorno>{
-------------------------------------------------
------------------------------------------------
------------------------------------------------
  //return en el caso de que nuestra funcion retorne valores
  }

*/

func saludar() {
	fmt.Println("Hola esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func main() {
	var usr string

	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)
}


func suma_resta (num1, num2 int) (int, int) {
	if num2<num1 {
		ResSuma := num1+num2
		ResResta := num2-num1
		return ResSuma, ResResta

	}else
}