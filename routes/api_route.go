package routes

import (
	"project-akhir-gdsc-backend/controllers"

	"github.com/gorilla/mux"
)

func AppRoutes(router *mux.Router) {
	router.HandleFunc("/api/imagestopdf", controllers.ConvertImagesToPDF).Methods("POST")
}
