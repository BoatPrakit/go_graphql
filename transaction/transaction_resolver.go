package transaction

import (
	"github.com/graphql-go/graphql"
)

type TransactionResolver struct {
	transaction *Storage
}

func NewTransactionResolver(transactionStorage *Storage) *TransactionResolver {
	return &TransactionResolver{
		transaction: transactionStorage,
	}
}

func (r *TransactionResolver) InsertTransaction(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["transaction"].(map[string]interface{})

	t := Transaction{
		BankAccountID: input["bankAccountId"].(string),
		Amount:        input["amount"].(int),
		Status:        input["status"].(int),
	}

	id, err := r.transaction.Create(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TransactionResolver) InsertMultipleTransaction(params graphql.ResolveParams) (interface{}, error) {
	transactions := params.Args["transactions"].([]interface{})
	results := []int64{}
	for _, t := range transactions {
		input := t.(map[string]interface{})
		transaction := Transaction{
			BankAccountID: input["bankAccountId"].(string),
			Amount:        input["amount"].(int),
			Status:        input["status"].(int),
		}
		id, err := r.transaction.Create(transaction)
		if err != nil {
			return 0, err
		}
		results = append(results, id)
	}

	return results, nil
}

func (r *TransactionResolver) GetAllTransaction(p graphql.ResolveParams) (interface{}, error) {
	transactions, err := r.transaction.FindAll()

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
