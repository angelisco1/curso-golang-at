package main

import "fmt"

type Numero int

type Direccion int

const (
	Norte Direccion = iota
	Sur
	Este
	Oeste
)

var toString = map[Direccion]string{
	Norte: "Norte",
	Sur:   "Sur",
	Este:  "Este",
	Oeste: "Oeste",
}

var toDireccion = map[string]Direccion{
	"Norte": Norte,
	"Sur":   Sur,
	"Este":  Este,
	"Oeste": Oeste,
}

func NewDireccion(s string) *Direccion {
	dir := toDireccion[s]
	return &dir
}

// func String(d *Direccion) string {
// 	return toString[*d]
// }
func (d Direccion) String() string {
	return toString[d]
}

type Brujula struct {
	Marca    string
	ApuntaAl *Direccion
}

func main() {
	// var num1 int = 2
	var num2 Numero = 3
	fmt.Println(DevuelveNumero(num2))
	// fmt.Println(DevuelveNumero(num1))

	// n := NewDireccion("Norte")
	// fmt.Println(*n)
	n := NewDireccion("Norte")
	fmt.Println(n)

	b := Brujula{
		Marca:    "Marca 1",
		ApuntaAl: NewDireccion("Sur"),
	}
	fmt.Println(b.ApuntaAl)

}

func DevuelveNumero(n Numero) Numero {
	return n
}
