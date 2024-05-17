package transaction

import "database/sql"

type Storage struct {
	db *sql.DB
}

func NewTransactionStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (ts *Storage) Create(t Transaction) (int64, error) {
	result, err := ts.db.Exec("INSERT INTO deposit_transaction ( bank_account_id, amount, status) VALUES ($1, $2, $3)", t.BankAccountID, t.Amount, t.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (ts *Storage) FindByID(id int64) (Transaction, error) {
	rows, err := ts.db.Query("SELECT * FROM deposit_transaction WHERE id = $1", id)
	if err != nil {
		return Transaction{}, err
	}

	transaction := Transaction{}
	for rows.Next() {
		err := rows.Scan(
			&transaction.ID,
			&transaction.BankAccountID,
			&transaction.Amount,
			&transaction.Status,
			&transaction.CreatedAt)
		if err != nil {
			return Transaction{}, err
		}
	}

	return transaction, nil

}

func (ts *Storage) FindAll() ([]Transaction, error) {
	rows, err := ts.db.Query("SELECT * FROM deposit_transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []Transaction{}
	for rows.Next() {
		t := Transaction{}
		err := rows.Scan(&t.ID,
			&t.BankAccountID,
			&t.Amount,
			&t.Status,
			&t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
