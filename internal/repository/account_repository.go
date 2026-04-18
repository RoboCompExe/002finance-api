package repository

import (
	"database/sql"
)

type AccountRepository struct {
	DB *sql.DB
}

func (r *AccountRepository) GetBalance(id string) (int64, error) {
	var balance int64
	err := r.DB.QueryRow("SELECT balance FROM accounts WHERE id=$1", id).Scan(&balance)
	return balance, err
}
