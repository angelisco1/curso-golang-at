package main

import (
	// "encoding/csv"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Persona struct {
	Nombre   string `csv:"nombre"`
	Apellido string `csv:"apellidos"`
	Email    string `csv:"email"`
	Rol      string `csv:"rol"`
}

func NewPersona(n, a, e, r string) *Persona {
	return &Persona{
		Nombre:   n,
		Apellido: a,
		Email:    e,
		Rol:      r,
	}
}

// func fn1() {

// 	defer fmt.Println(3)
// 	defer fmt.Println(4)
// }

func main() {

	// defer fmt.Println(1)
	// fn1()
	// defer fmt.Println(2)

	file, _ := os.Open("data/colegio.csv")
	defer file.Close()

	// filasCsv, _ := csv.NewReader(file).ReadAll()

	// var personas map[string]string
	var personas []Persona
	gocsv.UnmarshalFile(file, &personas)

	fmt.Println(personas)

	// fmt.Println(filasCsv)
	var alumnos, profesores []Persona

	// for _, fila := range filasCsv {
	// 	p := NewPersona(fila[0], fila[1], fila[2], fila[3])
	// 	if p.Rol == "profesor" {
	// 		profesores = append(profesores, *p)
	// 	} else {
	// 		alumnos = append(alumnos, *p)
	// 	}
	// }

	for _, p := range personas {
		if p.Rol == "profesor" {
			profesores = append(profesores, p)
		} else {
			alumnos = append(alumnos, p)
		}
	}

	// fmt.Println(profesores)
	// fmt.Println(alumnos)

	// crearArchivo(profesores, "data/profes.csv")
	// crearArchivo(alumnos, "data/alumnos.csv")
	crearArchivo(profesores, "data/profes-gocsv.csv")
	crearArchivo(alumnos, "data/alumnos-gocsv.csv")

}

func crearArchivo(datos []Persona, nombreArchivo string) {
	file, _ := os.Create(nombreArchivo)
	// writer := csv.NewWriter(file)

	// defer func() {
	// 	writer.Flush()
	// 	file.Close()
	// }()
	defer file.Close()

	gocsv.MarshalFile(datos, file)

	// for _, p := range datos {
	// 	// fila := fmt.Sprintf("%s,%s,%s,%s", p.Nombre, p.Apellido, p.Email, p.Rol)
	// 	writer.Write([]string{p.Nombre, p.Apellido, p.Email, p.Rol})
	// }

}
