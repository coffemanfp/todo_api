package database

import "fmt"

// RepositoryID is the key to use for the repositories hashmap.
type RepositoryID string

type Repositories map[RepositoryID]interface{}

func GetRepository[t any](repoMap Repositories, id RepositoryID) (repo t, err error) {
	repo, ok := repoMap[id].(t)
	if !ok {
		err = fmt.Errorf("missing repository: %s not found in repository map", id)
		return
	}
	return
}
