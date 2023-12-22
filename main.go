package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/AbdulRafayZia/graphql/db"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

// Joke represents the structure of a joke.
type Joke struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}

var jokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"joke": &graphql.Field{ // Update "content" to "joke"
				Type: graphql.String,
			},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"joke": &graphql.Field{
				Type: jokeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetRandomJoke(), nil
				},
			},
		},
	},
)

func GetRandomJoke() *Joke {
	var jokes []Joke

	db := db.OpenDB()

	// Retrieve jokes from the database and populate the 'jokes' slice.
	rows, err := db.Query("SELECT id, joke FROM jokes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var joke Joke
		err := rows.Scan(&joke.ID, &joke.Joke)
		if err != nil {
			log.Fatal(err)
		}
		jokes = append(jokes, joke)
	}
	if len(jokes) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(jokes))
	return &jokes[index]
}

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

func main() {
	// Connect to PostgreSQL database.

	// Set up the HTTP handler for GraphQL queries.
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        Schema,
			RequestString: r.URL.Query().Get("query"),
		})
		if len(result.Errors) > 0 {
			log.Printf("GraphQL query errors: %v", result.Errors)
		}

		json.NewEncoder(w).Encode(result)
	})

	// Start the server.
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
