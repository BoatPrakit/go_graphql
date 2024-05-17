package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/boatprakit/graphql/resolvers"
	"github.com/boatprakit/graphql/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Create db sqlited3
	db, err := sql.Open("sqlite3", "./contact.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Define GraphQL schema
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    createRootQuery(db),
		Mutation: createMutation(db),
	})
	// Create GraphQL handler

	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
		FormatErrorFn: func(err error) gqlerrors.FormattedError {
			slog.Error(err.Error())
			return gqlerrors.FormatError(err)
		},
	})

	// Serve GraphQL API at /graphql endpoint
	http.Handle("/graphql", graphqlHandler)

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:4000/graphql")
	http.ListenAndServe(":4000", nil)
}

// Define root query
func createRootQuery(db *sql.DB) *graphql.Object {
	r := resolvers.NewContactResolver(db)
	t := resolvers.NewTransactionResolver(db)
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					defer p.Context.Done()
					return "Hello, GraphQL!", nil
				},
			},
			"contact": &graphql.Field{
				Type:    graphql.NewList(types.ContactType),
				Resolve: r.Contacts,
			},
			"getContactById": &graphql.Field{
				Type: types.ContactType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: r.ContactById,
			},
			"transaction": &graphql.Field{
				Type:    graphql.NewList(types.TransactionType),
				Resolve: t.Transaction,
			},
		},
	})
	return rootQuery
}

func createMutation(db *sql.DB) *graphql.Object {
	r := resolvers.NewTransactionResolver(db)

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDepositTransaction": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"transaction": &graphql.ArgumentConfig{
						Type: types.TransactionInput,
					},
				},
				Resolve: r.InsertTransaction,
			},
			"createDepositTransactions": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
				Args: graphql.FieldConfigArgument{
					"transactions": &graphql.ArgumentConfig{
						Type: graphql.NewList(types.TransactionInput),
					},
				},
				Resolve: r.InsertMultipleTransaction,
			},
		},
	})
	return rootMutation
}
