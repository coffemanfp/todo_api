package database

import (
	"github.com/coffemanfp/todo/account"
)

// AUTH_REPOSITORY is the key to be used when creating the repositories hashmap.
const AUTH_REPOSITORY RepositoryID = "AUTH"

// AuthRepository defines the behaviors to be used by a AuthRepository implementation.
type AuthRepository interface {
	GetIdAndHashedPassword(account account.Account) (id int, hash string, err error)
	Register(accounut account.Account) (id int, err error)
}
