package psql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
)

type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository initializes a new account repository instance.
//
//	@param conn *PostgreSQLConnector: is the PostgreSQLConnector handler.
//	@return repo database.AccountRepository: is the final interface to keep
//	 the AccountRepository implementation.
//	@return err error: database connection error.
func NewAccountRepository(conn *PostgreSQLConnector) (repo database.AccountRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = AccountRepository{
		db: db,
	}
	return
}

func (ar AccountRepository) MatchCredentials(account account.Account) (id int, err error) {
	query := `
		select id from account where (nickname = $1 and password = $3) or (email = $2 and password = $3)
	`

	err = ar.db.QueryRow(query, account.Nickname, account.Email, account.Password).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
			return
		}
		err = fmt.Errorf("failed to get account credentials: %s", err)
	}
	return
}

func (ar AccountRepository) Register(account account.Account) (id int, err error) {
	query := `
		insert into
			account(name, last_name, nickname, email, password, created_at, updated_at)
		values
			($1, $2, $3, $4, $5, $6, $7)
		returning
			id
	`

	err = ar.db.QueryRow(query, account.Name, account.LastName, account.Nickname, account.Email, account.Password, account.CreatedAt, account.UpdatedAt).Scan(&id)
	if err != nil {
		err = fmt.Errorf("failed to insert record in account table: %s", err)
	}
	return
}
