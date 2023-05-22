package database

import "fmt"

// RepositoryID is the key to use for the repositories hashmap.
type RepositoryID string

type Repositories map[RepositoryID]interface{}

func GetRepository(repoMap Repositories, id RepositoryID) (repo interface{}, err error) {
	repo, ok := repoMap[id]
	if !ok {
		err = fmt.Errorf("missing repository: %s not found in repository map", id)
		return
	}
	return
}
