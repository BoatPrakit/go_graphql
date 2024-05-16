package models

type BankAccount struct {
	ID        int    `db:"id" json:"id"`
	AccountId string `db:"account_id" json:"accountId"`
	Name      string `db:"name" json:"name"`
	Balance   int    `db:"balance" json:"balance"`
}
