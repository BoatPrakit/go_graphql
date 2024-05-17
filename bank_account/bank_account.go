package bank_account

import "github.com/graphql-go/graphql"

type BankAccount struct {
	ID        int    `db:"id" json:"id"`
	AccountId string `db:"account_id" json:"accountId"`
	Name      string `db:"name" json:"name"`
	Balance   int    `db:"balance" json:"balance"`
}

var BankAccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankAccount",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"accountId": &graphql.Field{Type: graphql.String},
		"name":      &graphql.Field{Type: graphql.String},
		"balance":   &graphql.Field{Type: graphql.Int},
	},
},
)
