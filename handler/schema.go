package handler 

import (
	"github.com/graphql-go/graphql"
)

var jokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"Joke": &graphql.Field{ // Update "content" to "joke"
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

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)
