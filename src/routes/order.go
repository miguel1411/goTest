package routes

import (
	"github.com/gorilla/mux"
	"github.com/miguel1411/go-api-test/src/controller"
)

func Routesorders(r *mux.Router) {

	r.HandleFunc("/orders", controller.ClientAllGet).Methods("GET")
	r.HandleFunc("/orders/{name}/{lastname}/{email}", controller.ClientPost).Methods("POST")
	r.HandleFunc("/orders/{id}", controller.ClientDelete).Methods("PUT")

}
