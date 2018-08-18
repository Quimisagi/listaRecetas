package main

import(
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/gorilla/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "newpassword"
	dbname   = "recetas"
)

type Receta struct{
	ID string `json:"id"`
	Nombre string `json:"nombre"`
	Ingredientes []string `json:"ingredientes"`
	Instrucciones string `json:"instrucciones"`
}


func conectarDB()(*sql.DB, error){
	//Abriendo conexión con base de datos

	fmt.Println("Intentando abrir conexion")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado éxitosamente")

	return db, err

}



func crearReceta(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var receta Receta
	_ = json.NewDecoder(r.Body).Decode(&receta)

	db, err := conectarDB()

	sqlStatement1 := `
	INSERT INTO recetas (nombre_receta, instrucciones)
	VALUES ($1, $2)
	RETURNING id_receta;`
	sqlStatement2 := `
	INSERT INTO ingredientes (id_receta, nombre_ingrediente)
	VALUES ($1, $2)`

	err = db.QueryRow(sqlStatement1,receta.Nombre, receta.Instrucciones).Scan(&receta.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", receta.ID)

	ingredientes := receta.Ingredientes

	for _, item := range ingredientes{
		_, err = db.Exec(sqlStatement2, receta.ID, item)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}

	json.NewEncoder(w).Encode(receta)

}


func getRecetas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")


	db, err := conectarDB()

	sqlStatement1 := `SELECT id_receta, nombre_receta, instrucciones FROM recetas`
	sqlStatement2 := `SELECT nombre_ingrediente FROM ingredientes WHERE id_receta =$1;`

	var receta Receta
	var listaRecetas []Receta
	var ingrediente string

	rows, err := db.Query(sqlStatement1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&receta.ID, &receta.Nombre, &receta.Instrucciones)
		if err != nil {
			panic(err)
		}

		rows, err := db.Query(sqlStatement2, receta.ID)

		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&ingrediente)
			if err != nil {
				panic(err)
			}
			receta.Ingredientes = append(receta.Ingredientes, ingrediente)
		}

		listaRecetas = append(listaRecetas, receta)
		receta.Ingredientes = nil
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(listaRecetas)
}

func getReceta(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	params := mux.Vars(r)



	db, err := conectarDB()
	var receta Receta
	receta.ID = params["id"]
	var ingrediente string


	sqlStatement1 := `SELECT nombre_receta, instrucciones FROM recetas WHERE id_receta=$1;`
	sqlStatement2 := `SELECT nombre_ingrediente FROM ingredientes WHERE id_receta=$1;`

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement1, receta.ID)
	err = row.Scan(&receta.Nombre, &receta.Instrucciones)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		rows, err := db.Query(sqlStatement2, receta.ID)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {

			err = rows.Scan(&ingrediente)
			if err != nil {
				panic(err)
			}
			receta.Ingredientes = append(receta.Ingredientes, ingrediente)
		}
		json.NewEncoder(w).Encode(receta)
	default:
		panic(err)
	}

}
func actualizarReceta(w http.ResponseWriter, r *http.Request){


	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var receta Receta
	receta.ID = params["id"]

	_ = json.NewDecoder(r.Body).Decode(&receta)

	db, err := conectarDB()

	sqlStatement1 := `
	UPDATE recetas
	SET nombre_receta = $2, instrucciones = $3
	WHERE id_receta = $1;`
	_, err = db.Exec(sqlStatement1, receta.ID, receta.Nombre, receta.Instrucciones)
	if err != nil {
		panic(err)
	}

	sqlStatement2 := `
	DELETE FROM ingredientes
	WHERE id_receta = $1;`

	ingredientes := receta.Ingredientes

	_, err = db.Exec(sqlStatement2, receta.ID)

	sqlStatement3 := `
	INSERT INTO ingredientes (id_receta, nombre_ingrediente)
	VALUES ($1, $2)`

	for _, item := range ingredientes{
		_, err = db.Exec(sqlStatement3, receta.ID, item)
		if err != nil {
			panic(err)
		}
	}

	json.NewEncoder(w).Encode(receta)
}


func borrarReceta(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	id := params["id"]

	db, err := conectarDB()

	sqlStatement2 := `
	DELETE FROM ingredientes 
	WHERE id_receta = $1;`
	_, err = db.Exec(sqlStatement2, id)
	if err != nil {
		panic(err)
	}

	sqlStatement1 := `
	DELETE FROM recetas 
	WHERE id_receta = $1;`
	_, err = db.Exec(sqlStatement1, id)
	if err != nil {
		panic(err)
	}


}


func main() {


	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/api/listaRecetas", getRecetas).Methods("GET")
	router.HandleFunc("/api/listaRecetas/{id}", getReceta).Methods("GET")
	router.HandleFunc("/api/listaRecetas", crearReceta).Methods("POST")
	router.HandleFunc("/api/listaRecetas/{id}", actualizarReceta).Methods("PUT")
	router.HandleFunc("/api/listaRecetas/{id}", borrarReceta).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins) (router)))

}
