package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Schema
	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Author",
		Description: "An author",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(Author); ok {
						return author.Name, nil
					}
					return nil, nil
				},
			},
		},
	})

	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Book",
		Description: "A book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the book.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if book, ok := p.Source.(Book); ok {
						return book.Id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the book",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if book, ok := p.Source.(Book); ok {
						return book.Name, nil
					}
					return nil, nil
				},
			},
			"authors": &graphql.Field{
				Type: graphql.NewList(authorType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if book, ok := p.Source.(Book); ok {
						return book.Authors, nil
					}
					return nil, nil
				},
			},
		},
	})
	fields := graphql.Fields{
		"book": &graphql.Field{
			Type: bookType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var book = findOne()
				return book, nil
			},
		},
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var books = findAll()
				return books, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
