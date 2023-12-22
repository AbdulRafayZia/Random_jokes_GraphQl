package main

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/graphql/routes"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

// Joke represents the structure of a joke.

func main() {
	// Connect to PostgreSQL database.
	r := routes.Routes()
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)

	// Start the server
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
