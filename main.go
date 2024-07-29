package main

import (
	"log"
	"net/http"

	"ez-code-run/controllers"
	"ez-code-run/middlewares"
	"ez-code-run/models"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const directory = "public"

func main() {

	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB接続
	models.InitDB()

	// ルーティング
	mux := http.NewServeMux()

	// 静的ファイルのハンドリング
	mux.Handle("/", http.FileServer(http.Dir(directory)))

	// 認証関連のエンドポイント
	mux.HandleFunc("POST /api/auth/register", controllers.RegisterUser)
	mux.HandleFunc("POST /api/auth/login", controllers.LoginUser)
	mux.HandleFunc("GET /api/auth/user", middlewares.AuthMiddleware(controllers.GetUser))

	// Problem関連のエンドポイント
	mux.HandleFunc("GET /api/v1/problem", middlewares.AuthMiddleware(controllers.GetAllProblems))
	mux.HandleFunc("GET /api/v1/problem/{id}", middlewares.AuthMiddleware(controllers.GetProblem))

	// 提出関連のエンドポイント
	mux.HandleFunc("GET /api/v1/problem/{id}/submission", middlewares.AuthMiddleware(controllers.GetSubmissions))
	mux.HandleFunc("POST /api/v1/problem/{id}/submission", middlewares.AuthMiddleware(controllers.CreateSubmission))

	log.Println("Server is running on http://localhost:8080")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", "http://localhost:3002", "http://127.0.0.1:3002", "https://wis.yokoyama443.dev"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"*"},
		//AllowCredentials: true,
	})
	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
