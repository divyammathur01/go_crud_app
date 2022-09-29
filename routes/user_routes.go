package routes

import (
	"github.com/gorilla/mux"
	"github.com/rest_api_modular/controllers"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/createUser", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/getAllUsers", controllers.GetUsers()).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetSingleUser()).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")

}
