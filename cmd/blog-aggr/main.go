package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ortin779/blog-aggregator/handlers"
	"github.com/ortin779/blog-aggregator/middleware"
	"github.com/ortin779/blog-aggregator/store"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT Number not provided")
	}

	dbConnUrl := os.Getenv("DB_CONN_URL")
	if dbConnUrl == "" {
		log.Fatal("Database url not provided")
	}

	conn, err := sql.Open("postgres", dbConnUrl)

	if err != nil {
		log.Fatal("Error while connecting to database")
	}

	db := store.New(conn)
	userHandler := handlers.NewUserHandler(db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", handlers.RedinessHandler)
	mux.HandleFunc("GET /v1/err", handlers.ErrorHandler)

	mux.HandleFunc("POST /v1/users", userHandler.CreateUser)
	mux.HandleFunc("GET /v1/users", userHandler.GetUserByApikey)

	corsMux := middleware.Cors(mux)

	fmt.Println("Started server on Port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMux))
}
