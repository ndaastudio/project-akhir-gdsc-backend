package main

import (
	"log"
	"net/http"

	"project-akhir-gdsc-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.AppRoutes(router)

	log.Println("Server aktif!")
	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
