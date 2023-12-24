package database

import (
	"database/sql"
)

func InsertRequest(postgresdb *sql.DB, transaction Transaction) error {
	queryString := `
	INSERT INTO transaction (
		date_transaction, 
		amount, 
		account_id, 
		filename,
		input_timestamp) 
	VALUES ($1, $2, $3, $4, $5)`
	_, err := postgresdb.Exec(queryString,
		transaction.Date,
		transaction.Amount,
		transaction.IdAccount,
		transaction.Filename,
		transaction.Timestamp)

	if err != nil {
		return err
	}
	return nil
}
