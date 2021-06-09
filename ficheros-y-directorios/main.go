package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("archivos/archivo1.txt", os.O_APPEND|os.O_WRONLY, fs.FileMode(0700))
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("de cuyo nombre no quiero acordarme...")
	if err != nil {
		fmt.Println(err)
	}

	err = os.MkdirAll("archivos/data", fs.FileMode(0700))
	if err != nil {
		fmt.Println(err)
	}

	files, _ := os.ReadDir("archivos/imagenes")
	for i, file := range files {
		// fmt.Printf("%#v\n", file.parent)
		// fmt.Println(file.Name())
		rutaOrigen := fmt.Sprintf("archivos/imagenes/%s", file.Name())
		content, _ := ioutil.ReadFile(rutaOrigen)

		var extension string = strings.Split(file.Name(), ".")[1]
		rutaDestino := fmt.Sprintf("archivos/imagenes-renombradas/imagen-%d.%s", i, extension)

		// fmt.Println(rutaDestino)
		ioutil.WriteFile(rutaDestino, content, fs.FileMode(0700))
	}

}
