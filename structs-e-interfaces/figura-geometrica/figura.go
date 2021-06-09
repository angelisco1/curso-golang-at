package figura

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float32
}

type Rectangulo struct {
	Base   float32
	Altura float32
}

type Figura interface {
	Area() float32
}

type FiguraConF interface {
	Figura
	F()
}

func NewCirculo(radio float32) *Circulo {
	return &Circulo{
		Radio: radio,
	}
}

func NewRectangulo(base, altura float32) *Rectangulo {
	return &Rectangulo{
		Base:   base,
		Altura: altura,
	}
}

func (c *Circulo) Area() float32 {
	return math.Pi * float32(math.Pow(float64(c.Radio), 2.0))
}

func (r *Rectangulo) Area() float32 {
	return r.Base * r.Altura
}

func (r *Rectangulo) F() {
	fmt.Print("F")
}
