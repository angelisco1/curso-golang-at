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

	ar4Elementos := [...]int{1, 2, 3, 4}
	fmt.Println("4 Elementos: ", ar4Elementos)

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
	productosTienda4, _ := Pop(productosTienda2)
	fmt.Println(productosTienda4)

	datos := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7},
	}
	fmt.Println(datos[1][1])

	/* Maps */
	sugus := map[string]string{
		"piña":    "azul",
		"naranja": "naranja",
		"limon":   "amarillo",
		"fresa":   "",
	}

	fmt.Println(sugus)
	fmt.Println(sugus["limon"])

	if color, existe := sugus["fresa"]; existe {
		fmt.Println(color)
	} else {
		fmt.Println("La clave no existe")
	}

	sugus["fresa"] = "rojo"
	sugus["sandia"] = "rojo"
	fmt.Println(sugus)

	if _, existe := sugus["fresa"]; existe {
		fmt.Println("La clave existe")
	} else {
		fmt.Println("La clave no existe")
	}

	delete(sugus, "fresa")

	if _, existe := sugus["fresa"]; existe {
		fmt.Println("La clave existe")
	} else {
		fmt.Println("La clave no existe")
	}

	// fmt.Println(existe)

	notasPorTrimestre := map[string][]int{
		"q1": {1, 7, 9},
		"q2": {9, 7, 9},
		"q3": {9, 7, 10},
	}

	fmt.Println(notasPorTrimestre["q3"][1])

	slPrueba := []string{"a", "b", "c"}

	slPrueba, item := Pop(slPrueba)
	fmt.Println(item)
	fmt.Println(slPrueba)

	for k, v := range sugus {
		fmt.Printf("Clave %s - Valor %s\n", k, v)
	}

	for i, v := range notasInglesFinales {
		fmt.Printf("Pos %d - Valor %d\n", i, v)
	}

	for i, v := range productosTienda {
		fmt.Printf("Pos %d - Valor %s\n", i, v)
	}

	for i, filas := range datos {
		for j, columna := range filas {
			fmt.Printf("Fila %d Columna %d - %d\n", i, j, columna)
		}
	}

}

func suma(n1, n2 int) int {
	return n1 + n2
}

/* Pop */
// Función Pop que recibe un slice y devuelve el slice sin el último elemento
func Pop(slice []string) ([]string, string) {
	// return slice[0:len(slice)-1]
	return slice[:len(slice)-1], slice[len(slice)-1]
}
