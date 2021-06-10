package main

import (
	"fmt"
	"strings"
)

func ImprimirNum(n int) {
	fmt.Println(n)
}

func main() {

	// go ImprimirNum(1)
	// go ImprimirNum(2)
	// go ImprimirNum(3)

	// go func() {
	// 	fmt.Println("FIN")
	// }()

	// /*
	// 	Mas código
	// 	fn1()
	// 	fn2()
	// */

	// time.Sleep(time.Second)

	// mensajes := make(chan string, 2)

	// // Lectura: <-mensajes
	// // Escritura: mensajes <-

	// // go func() {
	// // 	time.Sleep(time.Second * 2)
	// // 	mensajes <- "Que tal"
	// // }()

	// go func() {
	// 	time.Sleep(time.Second * 4)
	// 	mensajes <- "Hola"
	// 	mensajes <- "Que tal"

	// 	// time.Sleep(time.Second * 3)
	// 	// mensajes <- "Que tal?"
	// }()

	// mensaje1 := <-mensajes
	// fmt.Println(mensaje1)
	// time.Sleep(time.Second * 5)

	// mensaje2 := <-mensajes
	// fmt.Println(mensaje2)

	datos := []string{"Alejandro", "Waldo", "Angelo", "Antonio", "Carlos", "Claudio", "Elisa", "Emma", "Francisco", "Jonathan", "José Luis", "Juan Antonio", "Juan Eduardo", "Julio", "Raul"}
	nombresPorRutina := 5
	numRutinas := len(datos) / nombresPorRutina

	nombresEncontrados := make(chan string)
	finGoroutines := make(chan bool)

	for i := 0; i < numRutinas; i++ {

		subSliceNombres := make([]string, nombresPorRutina)
		copy(subSliceNombres, datos[i*nombresPorRutina:i*nombresPorRutina+nombresPorRutina])

		go func(filtro string, datos []string, encontrados chan<- string, fin chan<- bool) {
			for _, nombre := range datos {
				if strings.Contains(nombre, filtro) {
					encontrados <- nombre
				}
			}
			fin <- true
		}("n", subSliceNombres, nombresEncontrados, finGoroutines)

	}

	var nombresConN []string

	i := 0
	for i < numRutinas {
		select {
		case nombre := <-nombresEncontrados:
			nombresConN = append(nombresConN, nombre)
			fmt.Println("Nombre encontrado: ", nombre)
		case <-finGoroutines:
			fmt.Println("Terminada una goroutine")
			i++
			// default
			// 	fmt.Println("Otra cosa")
		}
	}

	fmt.Println(nombresConN)
}
