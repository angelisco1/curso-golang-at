package main

import "fmt"

func main() {
	/*  Arrays  */
	var notasInglesFinales [3]int
	notasInglesFinales[0] = 7
	notasInglesFinales[1] = 6
	notasInglesFinales[2] = 8
	// notasInglesFinales[3] = 9

	var notasMatesFinales [3]int = [3]int{4, 7, 8}
	notasMatesFinales2 := [3]int{4, 7, 8}

	fmt.Println(notasInglesFinales, notasMatesFinales, notasMatesFinales2)

	longitudArray := len(notasMatesFinales2)
	fmt.Println(longitudArray)

	/*  Slices  */
	productosTienda := []string{"Tomate", "Maiz", "Cebolla", "Leche"}
	// productosTienda[4] = "Pavo"
	fmt.Println(cap(productosTienda))
	productosTienda = append(productosTienda, "Leche de coco")
	fmt.Println(cap(productosTienda))

	var dosUltimasNotas []int = notasInglesFinales[1:]
	fmt.Println(dosUltimasNotas)
	notasInglesFinales[2] = 4
	fmt.Println(dosUltimasNotas)
	dosUltimasNotas[0] = 9
	fmt.Println(notasInglesFinales)

	dosUltimasNotas = append(dosUltimasNotas, 10)
	dosUltimasNotas = append(dosUltimasNotas, 3)
	fmt.Println(notasInglesFinales)

	productosTienda2 := productosTienda
	// productosTienda2[0] = "Galletas principe"
	productosTienda3 := append(productosTienda2, "Galletas principe")

	fmt.Println(productosTienda)
	fmt.Println(productosTienda2)
	// productosTienda3[5] = "Pavo"
	// productosTienda3[6] = "Atún"
	// productosTienda3[7] = "Zumo de naranja"
	// productosTienda3[8] = "Este ya no se puede añadir"
	fmt.Println(cap(productosTienda3))

	s1 := make([]int, 4, 5)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1[3] = 4
	// s2 := append(s1, 3)

	s3 := s1
	fmt.Println(s1)
	fmt.Println(s3)

	// s2[0] = 0
	// fmt.Println(s1)
	// fmt.Println(s2)

	sl1 := []int{1, 2}
	sl2 := []int{3, 4, 2}
	// sl3 := append(sl1, 3, 4, 2)
	sl3 := append(sl1, sl2...)
	fmt.Println(sl3)

	fmt.Println(productosTienda2)
	productosTienda4 := Pop(productosTienda2)
	fmt.Println(productosTienda4)

}

func suma(n1, n2 int) int {
	return n1 + n2
}

/* Pop */
// Función Pop que recibe un slice y devuelve el slice sin el último elemento
func Pop(slice []string) []string {
	// return slice[0:len(slice)-1]
	return slice[:len(slice)-1]
}
