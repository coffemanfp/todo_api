package database

import "github.com/coffemanfp/todo/task"

const CATEGORY_REPOSITORY RepositoryID = "CATEGORY"

type CategoryRepository interface {
	CreateCategory(category task.Category) (id int, err error)
	CreateCategoryBinds(binds []*task.CategoryTaskBind) (err error)
	UpdateCategory(category task.Category) (err error)
	DeleteCategory(id int) (err error)
	GetSomeCategories(page, createdBy int) (ts []*task.Category, err error)
}
