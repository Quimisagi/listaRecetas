// Creador por David Crespo
/* Prueba expedientes.co:
Hacer una REST API que permita leer, crear, actualizar y borrar recetas de cocina. Implementarlo usando golang. */

package main

import(
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//---------------------------------------------------------------

// Estructura de la receta

type Receta struct{
	ID string `json:"id"`
	Nombre string `json:"nombre"`
	Ingredientes []string `json:"ingredientes"`
	Instrucciones string `json:"instrucciones"`
}

var cont = 2
var listaRecetas []Receta


//---------------------------------------------------------------

func getRecetas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listaRecetas)
}

func getReceta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	for _, item := range listaRecetas{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Receta{})
}

func crearReceta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var receta Receta
	_ = json.NewDecoder(r.Body).Decode(&receta)
	receta.ID =  strconv.Itoa(cont) // Mock ID
	cont = cont + 1
	listaRecetas =  append(listaRecetas, receta) 
	json.NewEncoder(w).Encode(receta)
}

func actualizarReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range listaRecetas {
		if item.ID == params["id"] {
			listaRecetas = append(listaRecetas[:index], listaRecetas[index+1:]...)
			var receta Receta
			_ = json.NewDecoder(r.Body).Decode(&receta)
			receta.ID = params["id"]
			listaRecetas = append(listaRecetas, receta)
			json.NewEncoder(w).Encode(receta)
			return
		}
	}
}
func borrarReceta(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range listaRecetas {
		if item.ID == params["id"] {
		listaRecetas = append(listaRecetas[:index], listaRecetas[index+1:]...)
		break	
		}
	}
	json.NewEncoder(w).Encode(listaRecetas)

}






//---------------------------------------------------------------




func main() {

	r := mux.NewRouter()

	listaRecetas = append(listaRecetas, Receta{ID: "1", Nombre: "Pollo frito", Ingredientes: []string{"Pollo","Miel"} , Instrucciones: "Fritar el pollo."})

	r.HandleFunc("/api/listaRecetas", getRecetas).Methods("GET")
	r.HandleFunc("/api/listaRecetas/{id}", getReceta).Methods("GET")
	r.HandleFunc("/api/listaRecetas", crearReceta).Methods("POST")
	r.HandleFunc("/api/listaRecetas/{id}", actualizarReceta).Methods("PUT")
	r.HandleFunc("/api/listaRecetas/{id}", borrarReceta).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}