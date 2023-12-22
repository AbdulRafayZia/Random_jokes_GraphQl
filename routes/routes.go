package routes

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/graphql/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Joke represents the structure of a joke.

	func Routes() *mux.Router {
		// Connect to PostgreSQL database.
		r := mux.NewRouter()

		// Set up the HTTP handler for GraphQL queries.
		r.HandleFunc("/graphql", handler.GraphQLHandler).Methods("POST")
		r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hello")
		}).Methods("GET")

		// Start the server.
		return r
	}
