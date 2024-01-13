// back/cmd/my-music-app/main.go

package main

import (
	"log"
	"github.com/abeh9197/my-music-app/back/internal/routes"
	"net/http"
)

func main() {
	mux := routes.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
