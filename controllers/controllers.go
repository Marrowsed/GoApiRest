package controllers

import (
	"api_rest/db"
	"api_rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllHistoricPeople(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json") custom way to do it
	var p []models.Person
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnId(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json") custom way to do it
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person
	db.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CreateHistoricPeople(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Person
	json.NewDecoder(r.Body).Decode(&newPerson)
	db.DB.Create(&newPerson)
	json.NewEncoder(w).Encode(newPerson)

}

func UpdateHistoricPeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person
	db.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	db.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func DeleteHistoricPeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person
	db.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)

}
