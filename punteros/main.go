package main

import "fmt"

func main() {

	texto := "Hola mundo"

	texto1 := texto
	fmt.Println(texto, texto1)
	texto1 = "Hola"
	fmt.Println(texto, texto1)
	var texto2 *string = &texto
	// texto2 := &texto
	// t := "Adios"
	// texto2 = &t
	*texto2 = "Adios"
	fmt.Println(texto, texto2)
	fmt.Println(texto, *texto2)

	nombres := []string{"Sara", "Angel", "Arya", "Robb"}
	// n := &nombres
	ultimoNombre := Pop(&nombres)
	fmt.Println(ultimoNombre)
	fmt.Println(nombres)

	primerNombre := Shift(&nombres)
	fmt.Println(primerNombre)
	fmt.Println(nombres)

	longitudNuevoSlice := Unshift(&nombres, "Rickon")
	fmt.Println(longitudNuevoSlice)
	fmt.Println(nombres)
	fmt.Printf("%p\n", &nombres)
	// fmt.Println(*n)
}

// (*sl[1:])
// *sl[1:]
// (*sl)[1:]

// slice1 := []{"a", "b", "c"}
// c := Pop(slice1)
// c -> "c" | slice1 -> ["a", "b"]
func Pop(slice *[]string) (item string) {
	item = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return
}

// slice1 := []{"a", "b", "c"}
// c := Shift(slice1)
// c -> "a" | slice1 -> ["b", "c"]
func Shift(slice *[]string) string {
	item := (*slice)[0]
	*slice = (*slice)[1:]
	return item
}

// slice1 := []{"a", "b", "c"}
// c := Unshift(slice1, "z")
// c -> "4" | slice1 -> ["z", "a", "b", "c"]
func Unshift(slice *[]string, item string) int {
	sl := []string{item}
	*slice = append(sl, (*slice)...)
	return len(*slice)
}

/* Pop */
// Función Pop que recibe un slice y devuelve el slice sin el último elemento
// func Pop(slice []string) ([]string, string) {
// 	// return slice[0:len(slice)-1]
// 	return slice[:len(slice)-1], slice[len(slice)-1]
// }
