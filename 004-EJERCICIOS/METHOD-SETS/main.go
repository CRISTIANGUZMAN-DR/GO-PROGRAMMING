package main

import "fmt"

type persona struct {
	nombre string
}

type animal struct {
	tipo string
}

type humano interface { // todos los tipos que reciban el metodo hablar implementan la interface humano, es decir son considerados de tipo humano
	hablar()
}

func (p *persona) hablar() { // receptor de tipo puntero
	fmt.Println("hola, soy", p.nombre)
}

func (a animal) hablar() {
	fmt.Println("Hola, soy un", a.tipo)
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	humano := persona{
		nombre: "Cristian",
	}

	mascota := animal{
		tipo: "Gato",
	}

	diAlgo(&humano) /*para poder utilizar un dato que sea de tipo persona, tengo que pasar un valor direccionado a un puntero, ya que quien esta recibiendo el metodo que es necesario
	implementar para ser parte de la interfaz, en este caso hablar() es un recepetor de tipo puntero*/
	fmt.Println("-----")
	diAlgo(&mascota) /*mientras que aca puedo darle un valor a la funcion de tipo pontero o un valor normal, ya que mascota que es de tipo animal, no esta siendo recibida por un puntero
	sino por un dato de valor normal */
}
