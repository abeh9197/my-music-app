// back/cmd/my-music-app/main.go

package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/abeh9197/my-music-app/back/internal/routes"
)

func main() {
	mux := routes.InitializeRoutes()

	// CORSミドルウェアの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // すべてのオリジンを許可（本番環境では具体的なオリジンを指定）,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true, // 必要に応じてクレデンシャルを許可
	})

	// CORSミドルウェアを使用してサーバーを起動
	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
