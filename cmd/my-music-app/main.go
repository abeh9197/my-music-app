package main

import (
	"log"
	"my-music-app/mod/internal/routes"
	"net/http"
)

func main() {
	mux := routes.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
