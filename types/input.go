package types

import "github.com/graphql-go/graphql"

var TransactionInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TransactionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"bankAccountId": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"amount":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"status":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
},
)
