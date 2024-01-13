package routes

import (
	"github.com/abeh9197/my-music-app/back/internal/handlers"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	return mux
}