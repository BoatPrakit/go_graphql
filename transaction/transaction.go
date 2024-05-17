package transaction

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Transaction struct {
	ID            int       `db:"id" json:"id"`
	BankAccountID string    `db:"bank_account_id" json:"bankAccountId"`
	Amount        int       `db:"amount" json:"amount"`
	Status        int       `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}

type TransactionStorage interface {
	Create(transaction Transaction) error
	FindByID(id int) (Transaction, error)
	FindAll() ([]Transaction, error)
}

var TransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Transaction",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"bankAccountId": &graphql.Field{Type: graphql.String},
		"amount":        &graphql.Field{Type: graphql.Int},
		"status":        &graphql.Field{Type: graphql.Int},
		"createdAt":     &graphql.Field{Type: graphql.DateTime},
	},
})

var TransactionInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TransactionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"bankAccountId": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"amount":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"status":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
},
)
