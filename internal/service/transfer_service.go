package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type TransferService struct {
	DB *sql.DB
}

func (s *TransferService) Transfer(ctx context.Context, from, to string, amount int64) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var balance int64

	err = tx.QueryRowContext(ctx,
		"SELECT balance FROM accounts WHERE id=$1 FOR UPDATE",
		from,
	).Scan(&balance)

	if err != nil {
		return err
	}

	if balance < amount {
		return errors.New("insufficient balance")
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE accounts SET balance = balance - $1 WHERE id=$2",
		amount, from,
	)

	_, err = tx.ExecContext(ctx,
		"UPDATE accounts SET balance = balance + $1 WHERE id=$2",
		amount, to,
	)

	_, err = tx.ExecContext(ctx,
		"INSERT INTO transactions(id, from_account, to_account, amount, status) VALUES ($1,$2,$3,$4,'SUCCESS')",
		uuid.New(), from, to, amount,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}
