package resolvers

import (
	"database/sql"

	"github.com/boatprakit/graphql/models"
	"github.com/graphql-go/graphql"
)

func NewTransactionResolver(db *sql.DB) *QueryResolver {
	return &QueryResolver{db}
}

func (r *QueryResolver) InsertTransaction(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["transaction"].(map[string]interface{})

	t := models.Transaction{
		BankAccountID: input["bankAccountId"].(string),
		Amount:        input["amount"].(int),
		Status:        input["status"].(int),
	}

	result, err := r.db.Exec("INSERT INTO deposit_transaction ( bank_account_id, amount, status) VALUES ($1, $2, $3)", t.BankAccountID, t.Amount, t.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
