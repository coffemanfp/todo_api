package database

import (
	"github.com/coffemanfp/todo/list"
)

const LIST_REPOSITORY RepositoryID = "LIST"

type ListRepository interface {
	CreateList(list list.List) (id int, err error)
	GetList(id int) (list list.List, err error)
	UpdateList(list list.List) (err error)
	DeleteList(id int) (err error)
}
