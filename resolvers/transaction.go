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

func (r *QueryResolver) InsertMultipleTransaction(params graphql.ResolveParams) (interface{}, error) {
	transactions := params.Args["transactions"].([]interface{})
	results := []int64{}
	for _, t := range transactions {
		input := t.(map[string]interface{})
		transaction := models.Transaction{
			BankAccountID: input["bankAccountId"].(string),
			Amount:        input["amount"].(int),
			Status:        input["status"].(int),
		}
		result, err := r.db.Exec("INSERT INTO deposit_transaction ( bank_account_id, amount, status) VALUES ($1, $2, $3)", transaction.BankAccountID, transaction.Amount, transaction.Status)
		if err != nil {
			return 0, err
		}
		id, _ := result.LastInsertId()
		results = append(results, id)
	}

	return results, nil
}

func (r *QueryResolver) Transaction(p graphql.ResolveParams) (interface{}, error) {
	rows, _ := r.db.Query("SELECT * FROM deposit_transaction")
	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.BankAccountID,
			&transaction.Amount,
			&transaction.Status,
			&transaction.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
