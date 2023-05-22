package database

import (
	"fmt"
)

// RepositoryID is the key to use for the repositories hashmap.
type RepositoryID string

// Database is the Database manager for connections and repository instancies.
type Database struct {
	Conn         DatabaseConnector
	Repositories map[RepositoryID]interface{}
}

// DatabaseConnector defines a database connector handler.
type DatabaseConnector interface {

	// Connect creates new connection of the database implementation.
	//  @return $1 error: database connection error
	Connect() error
}

func GetRepository(repoMap map[RepositoryID]interface{}, id RepositoryID) (repo interface{}, err error) {
	repo, ok := repoMap[id]
	if !ok {
		err = fmt.Errorf("missing repository: %s not found in repository map", id)
		return
	}
	return
}
