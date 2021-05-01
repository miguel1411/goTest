package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguel1411/go-api-test/src/database"
	"github.com/miguel1411/go-api-test/src/models"
)

//Obtener Todos Los Clientes
func OrderAllGet(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()

	var clients []models.Client
	db.Find(&clients)
	fmt.Println("{}", clients)
	json.NewEncoder(w).Encode(clients)
}

// Dar de alta un cliente
func OrderPost(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()

	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	email := vars["email"]

	result := db.Create(&models.Client{Nombre: name, Apellido: lastname, Correo: email})
	fmt.Println(result)
}

func OrderDelete(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	params := mux.Vars(r)
	id := params["id"]

	client := models.Client{}

	db.Delete(&client, id)
}
