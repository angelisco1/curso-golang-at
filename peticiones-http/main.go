package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Oferta struct {
	Nombre  string
	Empresa string
	Salario int
	Ciudad  string
}

const (
	URL = "https://ejemplos-dc1c1.firebaseio.com/curso-go/angel/ofertas"
	EXT = ".json"
)

func main() {
	// GetOfertas()
	// CrearOferta()
	DeleteOferta("-MbqU5hkuhwaU4eJbkYe")
}

func GetOfertas() {
	resp, _ := http.Get(URL + EXT)

	fmt.Println(resp)

	defer resp.Body.Close()

	var ofertas map[string]Oferta
	// json.Unmarshal(resp.Body.Read(), &ofertas)

	json.NewDecoder(resp.Body).Decode(&ofertas)

	fmt.Println(ofertas)

	for k, oferta := range ofertas {
		fmt.Printf("La oferta %s está asociada a la clave %s\n", oferta.Nombre, k)
	}

}

func CrearOferta() {

	oferta := Oferta{"Oferta 1", "Empresa 1", 21000, "Ciudad 1"}
	bytesOferta, _ := json.Marshal(oferta)

	req, _ := http.NewRequest(http.MethodPost, URL+EXT, bytes.NewReader(bytesOferta))

	// json.NewEncoder(req.Body).Encode(&oferta)

	client := http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp)
}

func DeleteOferta(id string) {
	url := fmt.Sprintf("%s/%s%s", URL, id, EXT)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	client := http.Client{
		// Timeout: time.Millisecond * 900,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
