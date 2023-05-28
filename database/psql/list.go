package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/list"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(conn *PostgreSQLConnector) (repo database.ListRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = ListRepository{
		db: db,
	}
	return
}

func (lr ListRepository) CreateList(l list.List) (id int, err error) {
	table := "list"
	query := fmt.Sprintf(`
		insert into
			%s(title, description, background_picture_url, created_by, created_at)
		values
			($1, $2, $3, $4, $5)
		returning
			id
	`, table)

	err = lr.db.QueryRow(query, l.Title, l.Description, l.BackgroundPictureURL, l.CreatedBy, l.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (lr ListRepository) GetList(id int) (l list.List, err error) {
	table := "list"
	query := fmt.Sprintf(`
		select
			id, title, description, background_picture_url, created_at, created_by
		from
			%s
		where
			id = $1
	`, table)

	err = lr.db.QueryRow(query, id).Scan(&l.ID, &l.Title, &l.Description, &l.BackgroundPictureURL, &l.CreatedAt, &l.CreatedBy)
	if err != nil {
		err = errorInRow(table, "get", err)
	}
	return
}

func (lr ListRepository) GetSomeLists(page, createdBy int) (ls []*list.List, err error) {
	table := "list"
	query := fmt.Sprintf(`
		select
			id, title, description, background_picture_url, created_at, created_by
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

	ls = make([]*list.List, 0)
	for rows.Next() {
		l := new(list.List)
		l.Title = new(string)
		l.Description = new(string)
		err = rows.Scan(&l.ID, &l.Title, &l.Description, &l.BackgroundPictureURL, &l.CreatedAt, &l.CreatedBy)
		if err != nil {
			err = errorInRow(table, "scan", err)
			ls = nil
			return
		}

		ls = append(ls, l)
	}
	err = rows.Err()
	if err != nil {
		ls = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}

func (lr ListRepository) UpdateList(l list.List) (err error) {
	table := "list"
	query := fmt.Sprintf(`
		update
			%s
		set
			title = coalesce($1, title),
			description = coalesce($2, description),
			background_picture_url = coalesce($3, background_picture_url)
		where
			id = $4
	`, table)

	_, err = lr.db.Exec(query, l.Title, l.Description, l.BackgroundPictureURL, l.ID)
	if err != nil {
		err = errorInRow(table, "update", err)
	}
	return
}

func (lr ListRepository) DeleteList(id int) (err error) {
	table := "list"
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
