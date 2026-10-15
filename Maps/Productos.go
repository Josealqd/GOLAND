package main

import (
	"fmt"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p Producto) ValorTotal() float64 {
	return p.Precio * float64(p.Cantidad)
}

func (p Producto) TieneStock() bool {
	return p.Cantidad > 0
}

func (p Producto) MostrarInfo() {
	estadoStock := "sin stock"
	if p.TieneStock() {
		estadoStock = "disponible"
	}

	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d\n", p.Cantidad)
	fmt.Printf("Valor total: $%.2f\n", p.ValorTotal())
	fmt.Printf("Estado de stock: %s\n", estadoStock)
	fmt.Println("--------------------------------")
}

func main() {
	fmt.Println("--- Inventario de Productos ---\n")

	producto1 := Producto{Nombre: "Laptop Vivobook", Precio: 850.50, Cantidad: 4}
	producto2 := Producto{Nombre: "Monitor Curvo", Precio: 210.00, Cantidad: 0}
	producto3 := Producto{Nombre: "Teclado Mecánico", Precio: 45.75, Cantidad: 15}

	producto1.MostrarInfo()
	producto2.MostrarInfo()
	producto3.MostrarInfo()
}