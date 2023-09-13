package psql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
)

type AuthRepository struct {
	db *sql.DB
}

// NewAuthRepository initializes a new auth repository instance.
//
//	@param conn *PostgreSQLConnector: is the PostgreSQLConnector handler.
//	@return repo database.AuthRepository: is the final interface to keep
//	 the AuthRepository implementation.
//	@return err error: database connection error.
func NewAuthRepository(conn *PostgreSQLConnector) (repo database.AuthRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = AuthRepository{
		db: db,
	}
	return
}

func (ar AuthRepository) GetIdAndHashedPassword(account account.Account) (id int, hashed string, err error) {
	table := "account"
	query := `
		select id, password from account where nickname = $1 or email = $2
	`

	err = ar.db.QueryRow(query, account.Nickname, account.Email).Scan(&id, &hashed)
	if err != nil {
		err = errorInRow(table, "get", err)
	}
	return
}

func (ar AuthRepository) Register(account account.Account) (id int, err error) {
	tx, err := ar.db.Begin()
	if err != nil {
		err = errorInTx("begin", err)
		return
	}
	defer tx.Rollback()

	table := "account"
	query := fmt.Sprintf(`
		insert into
			%s(name, last_name, nickname, email, password, created_at)
		values
			($1, $2, $3, $4, $5, $6)
		returning
			id
	`, table)

	err = tx.QueryRow(query, account.Name, account.LastName, account.Nickname, account.Email, account.Password, account.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
		return
	}

	now := time.Now()
	colors := []string{"blue", "red", "yellow", "black", "purple", "orange"}
	table = "category"
	query = fmt.Sprintf(`
		insert into
			%s(created_by, name, color, created_at)
		values
			($1, $2, $3, $4)
		returning
			id
	`, table)
	var lastCategoryID int

	for _, color := range colors {
		err = tx.QueryRow(query, id, "Category "+color, color, now).Scan(&lastCategoryID)
		if err != nil {
			err = errorInRow(table, "insert", err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		err = errorInTx("commit", err)
	}
	return
}
