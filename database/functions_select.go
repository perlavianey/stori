package database

import (
	"database/sql"
)

func GetAccountData(postgresdb *sql.DB, IdAccount string) (string, string, error) {
	queryString := `SELECT name,email from account where account_id = $1;`
	row := postgresdb.QueryRow(queryString, IdAccount)
	var name, email sql.NullString
	err := row.Scan(
		&name,
		&email,
	)
	if err != nil {
		return "", "", err
	}
	n := name.String
	e := email.String

	return n, e, nil
}
