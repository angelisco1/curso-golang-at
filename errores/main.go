package main

import (
	"errors"
	"fmt"

	"github.com/angelisco1/customerrors"
	errorspkg "github.com/pkg/errors"
)

func main() {

	res1, err := Sqrt(-4)
	if err != nil {
		fmt.Println(err)
		// return
		// log.Fatal()
		// os.Exit(-1)
	}

	fmt.Println(res1)

	// var res2 float32
	res2, err := SqrtConInfo(-4)
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Println(res2)

	res3, err := SqrtConCustomError(-4)
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Println(res3)

	_, err = SqrtConCustomError2(-4)
	if msg, ok := err.(*customerrors.ErrOpNumNegativo); ok {
		fmt.Println("-", msg)
	} else if msg, ok := err.(*customerrors.ErrOpNumPositivo); ok {
		fmt.Println("+", msg)
	}

	_, err = SqrtConTraza(-4)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errorspkg.Cause(err))
		fmt.Printf("%+v\n", err)
	}

}

func Sqrt(num float32) (float32, error) {
	if num < 0 {
		return 0, errors.New("no se puede hacer la raíz cuadrada de un número negativo")
	}
	return 10, nil
}

func SqrtConInfo(num float32) (float32, error) {
	if num < 0 {
		return 0, fmt.Errorf("no se puede hacer la raíz cuadrada de un número negativo (%f)", num)
	}
	return 10, nil
}

func SqrtConCustomError(num float32) (float32, error) {
	if num < 0 {
		return 0, customerrors.NewErrOpNumNegativo(num)
	}
	return 10, nil
}

func SqrtConCustomError2(num float32) (float32, error) {
	if num < 0 {
		return 0, customerrors.NewErrOpNumNegativo(num)
	}
	if num > 0 {
		return 0, customerrors.NewErrOpNumPositivo(num)
	}
	return 0, nil
}

func SqrtConTraza(num float32) (float32, error) {
	if num < 0 {
		return 0, errorspkg.Wrap(customerrors.NewErrOpNumNegativo(num), "SQRT -")
	}
	if num > 0 {
		return 0, errorspkg.Wrap(customerrors.NewErrOpNumPositivo(num), "SQRT +")
	}
	return 0, nil
}
