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
func ClientAllGet(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()

	var clients []models.Client
	db.Find(&clients)
	fmt.Println("{}", clients)
	json.NewEncoder(w).Encode(clients)
}

//Obtener Cliente por {ID}
func ClientGet(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	params := mux.Vars(r)
	id := params["id"]
	client := models.Client{}
	db.Find(&client, id)
	json.NewEncoder(w).Encode(client)
}

//Obtener Cliente por {ID}

func ClientGetEntity(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()
	params := mux.Vars(r)
	entidad := params["entidad"]
	var clients []models.Client
	db.Find(&clients)
	// fmt.Println("{}", clients)
	// json.NewEncoder(w).Encode(clients)

	client := models.Client{}
	db.Raw("SELECT * FROM clients WHERE entidad_federativa = ?", entidad).Scan(&client)
	// db.Find(&client, entidad_federativa)
	json.NewEncoder(w).Encode(client)
}

// Dar de alta un cliente
func ClientPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db := database.GetConnection()

	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	email := vars["email"]
	entidad := vars["entidad"]

	result := db.Create(&models.Client{Nombre: name, Apellido: lastname, Correo: email, EntidadFederativa: entidad})
	fmt.Println(result)
}

// Eliminar Cliente por {ID}
func ClientDelete(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")

	db := database.GetConnection()
	params := mux.Vars(r)
	id := params["id"]
	client := models.Client{}
	db.Delete(&client, id)

}

// Actualizar Cliente por {ID}
func ClientUpdate(w http.ResponseWriter, r *http.Request) {

	db := database.GetConnection()
	params := mux.Vars(r)
	id := params["id"]
	name := params["name"]
	lastname := params["lastname"]
	email := params["email"]

	client := models.Client{}
	db.Where("id=?", id).Find(&client)
	client.Nombre = name
	client.Correo = email
	client.Apellido = lastname
	db.Save(&client)
}
