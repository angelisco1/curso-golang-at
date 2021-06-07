package main

import (
	"fmt"
)

const (
	Num    = 4
	NumMax = 8
	NumMin = 2
)

func main() {

	/*
		Tipos de datos:
		- int8, int16, int32, int64
		- uint8, uint16, uint32, uint64
		- flota32, float64
		- string
		- bool
		- byte -> uint8
		- rune -> uint32
		- map, array, slice, struct
	*/

	var texto string
	var esVerdad bool
	var edad int
	edad = 37

	precio := 3.4
	precio = 8.5

	var precio2 float32 = 5.9

	// Empieza por mayúsculas -> es público
	// Empieza por minúsculas -> es privado

	// fmt.Println("Hola mundo")
	fmt.Println(texto)
	fmt.Println(esVerdad)
	fmt.Println(edad)
	fmt.Println(precio)
	fmt.Println(precio2)

	if esVerdad {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es mentira")
	}

	notaMates := 7
	if notaMates > 5 {
		fmt.Println("Aprobado")
	} else if notaMates == 5 {
		fmt.Println("Aprobado raspando")
	} else {
		fmt.Println("Suspenso")
	}

	if esPar := notaMates%2 == 0; esPar {
		fmt.Printf("Es par: %t\n", esPar)
	} else {
		fmt.Printf("Es par: %t\n", esPar)
	}
	// No funciona la siguiente línea porque esPar está en el ámbito del bloque if
	// fmt.Printf("Es par???: %t\n", esPar)

	switch notaMates {
	case 8, 9:
		fmt.Println("Notable")
	case 6, 7:
		fmt.Println("Bien")
	case 5:
		fmt.Println("Suficiente")
	default:
		fmt.Println("Insuficiente")
	}

	switch {
	case notaMates > 5:
		fmt.Println("Aprobado")
	case notaMates == 5:
		fmt.Println("Aprobado raspando")
	default:
		fmt.Println("Suspenso")
	}

	/*
		Paquete fmt
	*/
	texto = "mundo"

	fmt.Printf("El precio es %f\n", precio)
	fmt.Printf("El precio es %.2f\n", precio)
	fmt.Printf("Mi nota es %d\n", notaMates)
	fmt.Printf("El tipo de la nota es %T\n", notaMates)
	fmt.Printf("Hola %s\n", texto)

	holaMundo := fmt.Sprintf("--- Hola %s\n", texto)
	fmt.Printf(holaMundo)

	var bytesTexto []byte = []byte(texto)
	fmt.Printf("Bytes texto: %v\n", bytesTexto)
	fmt.Printf("Bytes texto: %v\n", string(bytesTexto))
	fmt.Printf(fmt.Sprintf("Bytes texto: %s\n", bytesTexto)) // Sprintf es más eficiente que el string()

	/*
		Operadores
		+, -, *, /, %, **, ++, --
		&&, ||, !
		<, >, <=, >=, ==, !=
		&, |, !
		+=, -=, *=, /=, %=
	*/

	/*
		Bucles
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		fmt.Println(i)
		if i == 15 {
			break
		}
		i++
	}

}
