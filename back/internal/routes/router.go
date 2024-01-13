package routes

import (
	"my-music-app/mod/internal/handlers"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	return mux
}