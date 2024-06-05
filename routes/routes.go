package routes

import (
	"api_rest/controllers"
	"api_rest/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContetTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/people", controllers.AllHistoricPeople).Methods("GET")
	r.HandleFunc("/api/people/{id}", controllers.ReturnId).Methods("GET")
	r.HandleFunc("/api/people/{id}", controllers.DeleteHistoricPeople).Methods("DELETE")
	r.HandleFunc("/api/people/{id}", controllers.UpdateHistoricPeople).Methods("PUT")
	r.HandleFunc("/api/people", controllers.CreateHistoricPeople).Methods("POST")
	log.Fatal(http.ListenAndServe(":3030", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
