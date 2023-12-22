package handler

import (
	"log"
	"math/rand"
	"time"

	"github.com/AbdulRafayZia/graphql/db"
	"github.com/AbdulRafayZia/graphql/models"
)

func GetRandomJoke() *models.Joke {
	var jokes []models.Joke

	db := db.OpenDB()

	// Retrieve jokes from the database and populate the 'jokes' slice.
	rows, err := db.Query("SELECT id, joke FROM jokes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var joke models.Joke
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
