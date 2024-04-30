package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ortin779/blog-aggregator/handlers"
	"github.com/ortin779/blog-aggregator/middleware"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", handlers.RedinessHandler)
	mux.HandleFunc("GET /v1/err", handlers.ErrorHandler)

	corsMux := middleware.Cors(mux)

	fmt.Println("Started server on Port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMux))
}
