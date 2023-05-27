package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(conn *PostgreSQLConnector) (repo database.TaskRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = TaskRepository{
		db: db,
	}
	return
}

func (lr TaskRepository) CreateTask(t task.Task) (id int, err error) {
	table := "task"
	query := fmt.Sprintf(`
		insert into
			%s(title, description, list_id, created_by, created_at)
		values
			($1, $2, $3, $4, $5)
		returning
			id
	`, table)
	err = lr.db.QueryRow(query, t.Title, t.Description, t.ListID, t.CreatedBy, t.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (lr TaskRepository) GetTask(id int) (t task.Task, err error) {
	table := "task"
	query := fmt.Sprintf(`
		select
			id, title, description, list_id, created_at, created_by
		from
			%s
		where
			id = $1
	`, table)

	t.Title = new(string)
	t.Description = new(string)
	err = lr.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.ListID, &t.CreatedAt, &t.CreatedBy)
	if err != nil {
		t = task.Task{}
		err = errorInRow(table, "get", err)
	}
	return
}

func (lr TaskRepository) GetSomeTasks(page, createdBy int) (ts []*task.Task, err error) {
	table := "task"
	query := fmt.Sprintf(`
		select
			id, title, description, list_id, created_at, created_by
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

	rows, err := lr.db.Query(query, createdBy, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	ts = make([]*task.Task, 0)
	for rows.Next() {
		t := new(task.Task)
		t.Title = new(string)
		t.Description = new(string)
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.ListID, &t.CreatedAt, &t.CreatedBy)
		if err != nil {
			err = errorInRow(table, "scan", err)
			ts = nil
			return
		}

		ts = append(ts, t)
	}
	err = rows.Err()
	if err != nil {
		ts = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}

func (lr TaskRepository) UpdateTask(t task.Task) (err error) {
	table := "task"
	query := fmt.Sprintf(`
		update
			%s
		set
			title = coalesce($1, title),
			description = coalesce($2, description),
			list_id = coalesce($3, list_id)
		where
			id = $4
	`, table)

	_, err = lr.db.Exec(query, t.Title, t.Description, t.ListID, t.ID)
	if err != nil {
		err = errorInRow(table, "update", err)
	}
	return
}

func (lr TaskRepository) DeleteTask(id int) (err error) {
	table := "task"
	query := fmt.Sprintf(`
		delete from
			%s
		where
			id = $1
	`, table)

	_, err = lr.db.Exec(query, id)
	if err != nil {
		err = errorInRow(table, "delete", err)
	}
	return
}