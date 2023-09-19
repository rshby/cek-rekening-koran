package coresystem

import (
	"context"
	"database/sql"
	"errors"
)

type AgroCoreRepository struct {
	Db *sql.DB
}

func NewAgroCoreRepository(db *sql.DB) *AgroCoreRepository {
	return &AgroCoreRepository{
		Db: db,
	}
}

// method to get data by nomor_rekening
func (a *AgroCoreRepository) GetByAccountNumber(ctx context.Context, tx *sql.Tx, accountNumber string) ([]CoreSystem, error) {
	sqlQuery, err := tx.PrepareContext(ctx, "select account_number, account_name, remark, dorc, amount, transaction_date, transaction_time from coresystem WHERE account_number=? ORDER BY transaction_date ASC, transaction_time ASC")
	if err != nil {
		return nil, err
	}

	rows, err := sqlQuery.QueryContext(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	var results []CoreSystem
	for rows.Next() {
		var result CoreSystem
		if err := rows.Scan(&result.AccountNumber, &result.AccountName, &result.Remark, &result.Dorc, &result.Amount, &result.TransactionDate, &result.TransactionTime); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, errors.New("record not found")
	}

	return results, nil
}
