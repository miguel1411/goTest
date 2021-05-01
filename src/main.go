package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/miguel1411/go-api-test/src/database"
	"github.com/miguel1411/go-api-test/src/models"
	"github.com/miguel1411/go-api-test/src/routes"
)

func main() {
	//database conexión
	database.GetConnection()

	models.DataClient()

	//ROUTER
	r := mux.NewRouter().StrictSlash(true)
	// rutas
	routes.Routesclients(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8083"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}
