package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest_api_modular/configs"
	"github.com/rest_api_modular/routes"
)

func main() {
	//router config
	router := mux.NewRouter()
	routes.UserRoutes(router)

	//database connection
	configs.ConnectDB()

	log.Fatal(http.ListenAndServe(":6000", router))

}
