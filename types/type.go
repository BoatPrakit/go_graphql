package types

import "github.com/graphql-go/graphql"

var ContactType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Contact",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"name":      &graphql.Field{Type: graphql.String},
		"firstName": &graphql.Field{Type: graphql.String},
		"lastName":  &graphql.Field{Type: graphql.String},
		"genderId":  &graphql.Field{Type: graphql.Int},
		"dob":       &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"phone":     &graphql.Field{Type: graphql.String},
		"address":   &graphql.Field{Type: graphql.String},
	},
},
)
