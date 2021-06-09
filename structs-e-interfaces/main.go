package main

import (
	"fmt"

	"github.com/angelisco1/figura-geometrica"
	"github.com/angelisco1/persona"
	"github.com/fatih/color"
)

// func (p persona.Persona) GetDatos() string {
// 	return p.Nombre + " " + p.Apellidos
// }

func main() {

	charly := persona.NewPersona("Charly", "")
	// fmt.Println(charly.Nombre, charly.Apellidos)
	charly.ImprimirPersona()

	// charly.Nombre = "Charles"
	// fmt.Println(charly.Nombre, charly.Apellidos)
	charly.ImprimirPersona()

	// charly.SetApellido("Falco", "Fernandez")
	// charly.ImprimirPersona()

	persona.SetApellido2(&charly, "Fernandez")
	charly.ImprimirPersona()

	// No se pueden añadir nuevas propiedades
	// charly.Email = "cfalco@gmail.com"

	color.Red("Hola mundo")

	profe := persona.NewProfesor("Charly", "Falco", "Colegio 1")
	// profe := persona.NewProfesor("Charly", "A", "Falco", "Colegio 1")
	fmt.Println(profe.Persona.Nombre)
	fmt.Println(profe.Nombre)
	fmt.Println(profe.Colegio)

	p1 := persona.Persona{
		Nombre: "Persona 1",
	}
	fmt.Println(p1)

	c := figura.NewCirculo(4.0)
	ImprimirFigura(c)

	// var r *figura.Rectangulo
	r := figura.NewRectangulo(2, 7)
	ImprimirFigura(r)
	ImprimirFiguraF(r)

}

func ImprimirFigura(f figura.Figura) {
	fmt.Println("El área es: ", f.Area())
}

func ImprimirFiguraF(f figura.FiguraConF) {
	fmt.Println("El área es: ", f.Area())
}
