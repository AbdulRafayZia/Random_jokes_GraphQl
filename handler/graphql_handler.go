package handler

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})
	h.ServeHTTP(w, r)
}
