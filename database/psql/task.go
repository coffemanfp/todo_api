package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/search"
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

func (tr TaskRepository) CreateTask(t task.Task) (id int, err error) {
	table := "task"
	query := fmt.Sprintf(`
		insert into
			%s(title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_by, created_at)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		returning
			id
	`, table)
	err = tr.db.QueryRow(query, t.Title, t.Description, &t.ListID, t.Reminder, t.DueDate, t.Repeat, t.IsDone, t.IsAddedToMyDay, t.IsImportant, t.CreatedBy, t.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (tr TaskRepository) GetTask(id int) (t task.Task, err error) {
	table := "task"
	query := fmt.Sprintf(`
		select
			id, title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_at, created_by
		from
			%s
		where
			id = $1
	`, table)

	t.Title = new(string)
	t.Description = new(string)
	err = tr.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.ListID, &t.Reminder, &t.DueDate, &t.Repeat, &t.IsDone, &t.IsAddedToMyDay, &t.IsImportant, &t.CreatedAt, &t.CreatedBy)
	if err != nil {
		t = task.Task{}
		err = errorInRow(table, "get", err)
		return
	}

	query = `
		select
			c.id, name, color, created_at
		from
			category c
		inner join
			task_category tc on c.id = tc.category_id
		where
			task_id = $1
	`

	rows, err := tr.db.Query(query, t.ID)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	cs := make([]*task.Category, 0)
	for rows.Next() {
		c := new(task.Category)
		c.Name = new(string)
		c.Color = new(string)
		err = rows.Scan(&c.ID, &c.Name, &c.Color, &c.CreatedAt)
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
		return
	}

	t.Categories = cs
	return
}

func (tr TaskRepository) GetSomeTasks(page, listID, createdBy int) (ts []*task.Task, err error) {
	table := "task"

	var conditionalWhereStr string

	fmt.Println("listID:", listID)
	if listID != 0 {
		conditionalWhereStr = fmt.Sprintf("and list_id = %d", listID)
	}

	query := fmt.Sprintf(`
		select
			id, title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_at, created_by
		from
			%s
		where
			created_by = $1 %s
		limit
			$2
		offset
			$3
	`, table, conditionalWhereStr)

	limit, offset := parsePagination(page)

	rows, err := tr.db.Query(query, createdBy, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	ts = make([]*task.Task, 0)
	for rows.Next() {
		t := new(task.Task)
		t.Title = new(string)
		t.Description = new(string)
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.ListID, &t.Reminder, &t.DueDate, &t.Repeat, &t.IsDone, &t.IsAddedToMyDay, &t.IsImportant, &t.CreatedAt, &t.CreatedBy)
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

func (tr TaskRepository) Search(search search.Search) (ts []*task.Task, err error) {
	table := "task"
	query := fmt.Sprintf(`
        select
            id, title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_at, created_by
        from
            %s
		where
			($1::boolean is null or ($1::boolean is not null and is_done = $1)) and
			($2::boolean is null or ($2::boolean is not null and is_added_to_my_day = $2)) and
			($3::boolean is null or ($3::boolean is not null and is_important = $3)) and
			($4::boolean is null or ($4::boolean is not null and nullif(due_date, '0001-01-01 00:00:00') is not null = $4)) and
			($5::boolean is null or 
				($5::boolean is not null and nullif(due_date, '0001-01-01 00:00:00') is not null = $5 and (due_date - '3 days'::interval < now()) = $5)) and
			($6::text is null or ($6::text is not null and lower(title::text) LIKE CONCAT('%%',lower($6::text),'%%'))) and
			created_by = $7 
		`, table)
	rows, err := tr.db.Query(query, search.IsDone, search.IsAddedToMyDay, search.IsImportant, search.HasDueDate, search.ExpireSoon, search.Title, search.ClientID)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	ts = make([]*task.Task, 0)
	for rows.Next() {
		t := new(task.Task)
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.ListID, &t.Reminder, &t.DueDate, &t.Repeat, &t.IsDone, &t.IsAddedToMyDay, &t.IsImportant, &t.CreatedAt, &t.CreatedBy)
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

func (tr TaskRepository) UpdateTask(t task.Task) (err error) {
	tx, err := tr.db.Begin()
	if err != nil {
		err = errorInTx("begin", err)
		return
	}
	defer tx.Rollback()

	table := "task"
	query := fmt.Sprintf(`
		update
			%s
		set
			title = coalesce($1, title),
			description = coalesce($2, description),
			list_id = coalesce($3, list_id),
			reminder = coalesce($4, reminder),
			due_date = coalesce($5, due_date),
			repeat = coalesce($6, repeat),
			is_done = coalesce($7, is_done),
			is_added_to_my_day = coalesce($8, is_added_to_my_day),
			is_important = coalesce($9, is_important)
		where
			id = $10
	`, table)

	_, err = tx.Exec(query, t.Title, t.Description, t.ListID, t.Reminder, t.DueDate, t.Repeat, t.IsDone, t.IsAddedToMyDay, t.IsImportant, t.ID)
	if err != nil {
		err = errorInRow(table, "update", err)
		return
	}

	table = "task_category"
	query = fmt.Sprintf(`
		delete from
			%s
		where
			task_id = $1
	`, table)

	_, err = tr.db.Exec(query, t.ID)
	if err != nil {
		err = errorInRow(table, "delete", err)
		return
	}

	table = "task_category"
	query = fmt.Sprintf(`
		insert into
			%s(task_id, category_id)
		values
			($1, $2)
		returning
			id
	`, table)

	for _, c := range t.Categories {
		_, err = tx.Exec(query, t.ID, c.ID)
		if err != nil {
			err = errorInRow(table, "insert", err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		err = errorInTx("commit", err)
	}
	return
}

func (tr TaskRepository) DeleteTask(id int) (err error) {
	table := "task"
	query := fmt.Sprintf(`
		delete from
			%s
		where
			id = $1
	`, table)

	_, err = tr.db.Exec(query, id)
	if err != nil {
		err = errorInRow(table, "delete", err)
	}
	return
}
