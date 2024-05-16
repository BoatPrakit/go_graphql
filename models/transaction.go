package models

import "time"

type Transaction struct {
	ID            int       `db:"id" json:"id"`
	BankAccountID int       `db:"bank_account_id" json:"bankAccountId"`
	Amount        int       `db:"amount" json:"amount"`
	Status        int       `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}
