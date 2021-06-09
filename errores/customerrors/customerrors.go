package customerrors

import "fmt"

type ErrOpNumNegativo struct {
	Num float32
}

type ErrOpNumPositivo struct {
	Num float32
}

func NewErrOpNumNegativo(num float32) *ErrOpNumNegativo {
	return &ErrOpNumNegativo{
		Num: num,
	}
}

func (err ErrOpNumNegativo) Error() string {
	return fmt.Sprintf("no se puede hacer la raíz cuadrada de un número negativo (%.2f)", err.Num)
}

func NewErrOpNumPositivo(num float32) *ErrOpNumPositivo {
	return &ErrOpNumPositivo{
		Num: num,
	}
}

func (err ErrOpNumPositivo) Error() string {
	return fmt.Sprintf("no se puede hacer la raíz cuadrada de un número positivo (%.2f)", err.Num)
}
