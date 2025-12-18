package main

import (
	"folio/internal/auth"
	"folio/internal/db"
	"log"
	"net/http"
)

func main() {
	jwtSecret := "mysecretkey"

	dbConn, err := db.NewSqLiteDb("./data/folio.db")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := db.NewUserRepo(dbConn)

	authService := auth.NewAuthService(userRepo, jwtSecret)

	authHandler := &auth.AuthHandler{
		AuthService: authService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
