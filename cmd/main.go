package main

import (
	"log"
	"net/http"

	"github.com/rcgc/person-api-http-net/handler"
	"github.com/rcgc/person-api-http-net/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}