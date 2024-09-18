package main

import (
	"fmt"
	"go-programming/003-PAQUETES/producto"
)

func main() {

	fmt.Println("El vendedor se llama Cristian")

	producto.ConseguirProductos()
	producto.ObtenerDescuentos()

}
