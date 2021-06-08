package persona

import (
	"fmt"
)

type Persona struct {
	Nombre     string
	Apellidos  string
	Trabajando bool
}

type Profesor struct {
	Colegio string
	// Nombre  string
	Persona
}

// func NewProfesor(nombre1, nombre2, apellidos, colegio string) Profesor {
func NewProfesor(nombre1, apellidos, colegio string) Profesor {
	return Profesor{
		// Nombre:  nombre2,
		Colegio: colegio,
		Persona: Persona{
			Nombre:    nombre1,
			Apellidos: apellidos,
		},
	}
}

func (p Persona) ImprimirPersona() {
	fmt.Println(p.Nombre, p.Apellidos)
}

func NewPersona(nombre, apellidos string) Persona {
	return Persona{
		Nombre: nombre,
	}
}

func (p *Persona) SetApellido(apellido1, apellido2 string) {
	apellidos := fmt.Sprintf("%s %s", apellido1, apellido2)
	// apellidos := apellido1 + " " + apellido2
	// apellidos := strings.Join([]string{apellido1, apellido2}, " ")
	fmt.Println(apellidos)
	// (*p).apellidos = apellidos
	fmt.Printf("%p\n", p)
	// *p = Persona{
	// 	Nombre:    "a",
	// 	apellidos: "b",
	// }
	p.Apellidos = apellidos
}

func SetApellido2(p *Persona, apellido1 string) {
	p.Apellidos = apellido1
}
