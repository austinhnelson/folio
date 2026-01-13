package main

import (
	"folio/internal/auth"
	"folio/internal/db"
	"folio/internal/middleware"
	"log"
	"net/http"
)

func main() {
	dbConn, err := db.NewSqLiteDb("./data/folio.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.RunMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	userRepo := db.NewUserRepo(dbConn)

	authService := auth.NewAuthService(userRepo, "secret-jwt-key")

	authHandler := &auth.AuthHandler{
		AuthService: authService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	handler := middleware.CorsMiddleware(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
