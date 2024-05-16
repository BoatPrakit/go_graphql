package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

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
		Query: createRootQuery(db),
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
				Type: graphql.NewList(types.ContactType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, _ := db.Query("SELECT * FROM contact")
					var contacts []Contact
					for rows.Next() {
						var contact Contact
						err := rows.Scan(&contact.ID, &contact.Name, &contact.FirstName, &contact.LastName, &contact.GenderID, &contact.DOB, &contact.Email, &contact.Phone, &contact.Address, &contact.PhotoPath, &contact.CreatedAt, &contact.CreatedBy)

						if err != nil {
							return nil, err
						}
						contacts = append(contacts, contact)
					}
					return contacts, nil
				},
			},
		},
	})
	return rootQuery
}

type Contact struct {
	ID        int64       `db:"contact_id" json:"contactId"`
	Name      string      `db:"name" json:"name"`
	FirstName string      `db:"first_name" json:"firstName"`
	LastName  string      `db:"last_name" json:"lastName"`
	GenderID  int         `db:"gender_id" json:"genderId"`
	DOB       interface{} `db:"dob" json:"dob"`
	Email     string      `db:"email" json:"email"`
	Phone     string      `db:"phone" json:"phone"`
	Address   string      `db:"address" json:"address"`
	PhotoPath string      `db:"photo_path" json:"photoPath"`
	CreatedAt string      `db:"created_at" json:"createdAt"`
	CreatedBy string      `db:"created_by" json:"createdBy"`
}
