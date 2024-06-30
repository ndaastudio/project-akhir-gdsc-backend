package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server aktif!")
	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
