package database

import (
	"fmt"

	"github.com/coffemanfp/todo/account"
)

// ACCOUNT_REPOSITORY is the key to be used when creating the repositories hashmap.
const ACCOUNT_REPOSITORY RepositoryID = "ACCOUNT"

// GetAccountRepository gets the AccountRepository instance inside the repositories hashmap.
//
//		@param repoMap Repositories: repositories hashmap.
//		@return repo AccountRepository: found AccountRepository instance.
//	 @return err error: missing or invalid repository instance error.
func GetAccountRepository(repoMap Repositories) (repo AccountRepository, err error) {
	repoI, err := GetRepository(repoMap, ACCOUNT_REPOSITORY)
	repo, ok := repoI.(AccountRepository)
	if !ok {
		err = fmt.Errorf("invalid repository value: %s has a invalid %s repository handler", ACCOUNT_REPOSITORY, ACCOUNT_REPOSITORY)
	}
	return
}

// AccountRepository defines the behaviors to be used by a AccountRepository implementation.
type AccountRepository interface {
	MatchCredentials(account account.Account) (id int, err error)
	Register(accounut account.Account) (id int, err error)
}
