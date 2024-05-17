package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/boatprakit/graphql/auth"
	"github.com/boatprakit/graphql/contact"
	"github.com/boatprakit/graphql/transaction"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"
	_ "github.com/mattn/go-sqlite3"
)

type rootQuery struct {
	contactStorage     *contact.Storage
	transactionStorage *transaction.Storage
}
type rootMutation struct {
	transactionStorage *transaction.Storage
	authStorage        *auth.Storage
}

func main() {

	// Create db sqlited3
	db, err := sql.Open("sqlite3", "./contact.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create contact storage
	contactStorage := contact.NewStorage(db)
	transactionStorage := transaction.NewTransactionStorage(db)
	authStorage := auth.NewStorage(db)

	// Create root query
	rootQuery := rootQuery{
		contactStorage:     contactStorage,
		transactionStorage: transactionStorage,
	}

	rootMutation := rootMutation{
		transactionStorage: transactionStorage,
		authStorage:        authStorage,
	}

	// Define GraphQL schema
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    createRootQuery(rootQuery),
		Mutation: createMutation(rootMutation),
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
func createRootQuery(rq rootQuery) *graphql.Object {
	r := contact.NewContactResolver(rq.contactStorage)
	t := transaction.NewTransactionResolver(rq.transactionStorage)
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
				Type:    graphql.NewList(contact.ContactType),
				Resolve: r.Contacts,
			},
			"getContactById": &graphql.Field{
				Type: contact.ContactType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: r.ContactById,
			},
			"transaction": &graphql.Field{
				Type:    graphql.NewList(transaction.TransactionType),
				Resolve: t.GetAllTransaction,
			},
		},
	})
	return rootQuery
}

func createMutation(rm rootMutation) *graphql.Object {
	r := transaction.NewTransactionResolver(rm.transactionStorage)
	a := auth.NewResolver(rm.authStorage)

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDepositTransaction": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"transaction": &graphql.ArgumentConfig{
						Type: transaction.TransactionInput,
					},
				},
				Resolve: r.InsertTransaction,
			},
			"createDepositTransactions": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
				Args: graphql.FieldConfigArgument{
					"transactions": &graphql.ArgumentConfig{
						Type: graphql.NewList(transaction.TransactionInput),
					},
				},
				Resolve: r.InsertMultipleTransaction,
			},
			"login": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "LoginResponse",
					Fields: graphql.Fields{
						"token": &graphql.Field{Type: graphql.String},
					},
				}),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: a.Login,
			},
			"register": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"user": &graphql.ArgumentConfig{
						Type: graphql.NewInputObject(graphql.InputObjectConfig{
							Name: "RegisterInput",
							Fields: graphql.InputObjectConfigFieldMap{
								"name":     &graphql.InputObjectFieldConfig{Type: graphql.String},
								"email":    &graphql.InputObjectFieldConfig{Type: graphql.String},
								"password": &graphql.InputObjectFieldConfig{Type: graphql.String},
							},
						}),
					},
				},
				Resolve: a.Register,
			},
		},
	})
	return rootMutation
}
