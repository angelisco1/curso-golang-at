package producto

import (
	"fmt"

	"github.com/google/uuid"
)

type Producto struct {
	Id     string  `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float32 `json:"precio"`
}

var productos []Producto = []Producto{
	NewProducto("Producto 1", 23.78),
	NewProducto("Producto 2", 3.95),
	NewProducto("Producto 3", 10.5),
	NewProducto("Producto 4", 20.01),
}

func NewProducto(nombre string, precio float32) Producto {
	id := uuid.New().String()
	return Producto{id, nombre, precio}
}

func GetProductos() []Producto {
	return productos
}

func GetProductoById(id string) (*Producto, error) {
	for _, p := range productos {
		if p.Id == id {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("el producto con id %s no existe", id)
}

func CreateProducto(nombre string, precio float32) (*Producto, error) {
	for _, p := range productos {
		if p.Nombre == nombre {
			return nil, fmt.Errorf("el producto que intentas guardar ya existe")
		}
	}

	p := NewProducto(nombre, precio)

	productos = append(productos, p)
	return &p, nil
}
