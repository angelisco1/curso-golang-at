package main

import (
	"encoding/json"
	"net/http"

	producto "github.com/angelisco1/productos"
	"github.com/gorilla/mux"
)

func HandleGetProductos(w http.ResponseWriter, r *http.Request) {
	productos := producto.GetProductos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productos)
}

func HandleGetProducto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	producto, err := producto.GetProductoById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(producto)
}

func HandlePostProductos(w http.ResponseWriter, r *http.Request) {

	var datos *producto.Producto
	json.NewDecoder(r.Body).Decode(&datos)

	producto, err := producto.CreateProducto(datos.Nombre, datos.Precio)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(producto)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/productos", HandleGetProductos).Methods(http.MethodGet)
	r.HandleFunc("/productos/{id}", HandleGetProducto).Methods(http.MethodGet)
	r.HandleFunc("/productos", HandlePostProductos).Methods(http.MethodPost)
	// r.HandleFunc("/productos/id1", HandleGetProductos)

	http.ListenAndServe(":8080", r)
}
