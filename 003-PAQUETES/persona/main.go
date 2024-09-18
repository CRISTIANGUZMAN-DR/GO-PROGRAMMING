package main

import (
	"fmt"

	"github.com/GO-PROGRAMMING/003-PAQUETES/producto"
)

func main() {
	fmt.Println("El vendedor se llama Cristian")

	producto.ObtenerDescuentos()
	producto.ConseguirProductos()
}
