package database

// ACCOUNT_REPOSITORY is the key to be used when creating the repositories hashmap.
const ACCOUNT_REPOSITORY RepositoryID = "ACCOUNT"

// GetAccountRepository gets the AccountRepository instance inside the repositories hashmap.
//
//		@param repoMap map[RepositoryID]interface{}: repositories hashmap.
//		@return repo AccountRepository: found AccountRepository instance.
//	 @return err error: missing or invalid repository instance error.
func GetAccountRepository(repoMap map[RepositoryID]interface{}) (repo AccountRepository, err error) {
	return GetRepository(repoMap, ACCOUNT_REPOSITORY)
}

// AccountRepository defines the behaviors to be used by a AccountRepository implementation.
type AccountRepository interface{}
