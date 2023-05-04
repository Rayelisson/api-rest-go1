package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api-rest-1/controller"
	"api-rest-1/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.CriarUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controller.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
