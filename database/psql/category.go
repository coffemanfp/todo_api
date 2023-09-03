package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(conn *PostgreSQLConnector) (repo database.CategoryRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = CategoryRepository{
		db: db,
	}
	return
}

func (cr CategoryRepository) CreateCategory(c task.Category) (id int, err error) {
	table := "category"
	query := fmt.Sprintf(`
		insert into
			%s(created_by, name, color, created_at)
		values
			($1, $2, $3, $4)
		returning
			id
	`, table)
	err = cr.db.QueryRow(query, c.CreatedBy, c.Name, c.Color, c.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (cr CategoryRepository) CreateCategoryBind(taskID, categoryID int) (id int, err error) {
	table := "task_category"
	query := fmt.Sprintf(`
		insert into
			%s(task_id, category_id)
		values
			($1, $2)
		returning
			id
	`, table)
	err = cr.db.QueryRow(query, taskID, categoryID).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (cr CategoryRepository) GetSomeCategories(page, createdBy int) (cs []*task.Category, err error) {
	table := "category"
	query := fmt.Sprintf(`
		select
			id, created_by, name, color, created_at
		from
			%s
		where
			created_by = $1
		limit
			$2
		offset
			$3
	`, table)

	limit, offset := parsePagination(page)

	rows, err := cr.db.Query(query, createdBy, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	cs = make([]*task.Category, 0)
	for rows.Next() {
		c := new(task.Category)
		err = rows.Scan(&c.ID, &c.CreatedBy, &c.Name, &c.Color, &c.CreatedAt)
		if err != nil {
			err = errorInRow(table, "scan", err)
			cs = nil
			return
		}

		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		cs = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}

func (cr CategoryRepository) UpdateCategory(c task.Category) (err error) {
	table := "category"
	query := fmt.Sprintf(`
		update
			%s
		set
			name = coalesce($1::varchar, name),
			color = coalesce($2::varchar, color)
		where
			id = $3::integer
	`, table)

	_, err = cr.db.Exec(query, c.Name, c.Color, c.ID)
	if err != nil {
		err = errorInRow(table, "update", err)
	}
	return
}

func (cr CategoryRepository) DeleteCategory(id int) (err error) {
	table := "category"
	query := fmt.Sprintf(`
		delete from
			%s
		where
			id = $1
	`, table)

	_, err = cr.db.Exec(query, id)
	if err != nil {
		err = errorInRow(table, "delete", err)
	}
	return
}
