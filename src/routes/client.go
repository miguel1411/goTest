package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguel1411/go-api-test/src/controller"
)

func Routesclients(r *mux.Router) {
	r.HandleFunc("/clients/{id}", controller.ClientGet).Methods("GET")
	r.HandleFunc("/clients/{entidad}", controller.ClientGetEntity).Methods("GET")
	r.HandleFunc("/clients", controller.ClientAllGet).Methods("GET")
	r.HandleFunc("/clients/{name}/{lastname}/{email}/{entidad}", controller.ClientPost).Methods("POST", http.MethodPost)
	r.HandleFunc("/clients/{id}", controller.ClientDelete).Methods("PUT")                           // Eliminar
	r.HandleFunc("/clients/{id}/{name}/{lastname}/{email}", controller.ClientUpdate).Methods("PUT") // Actualizar
}
