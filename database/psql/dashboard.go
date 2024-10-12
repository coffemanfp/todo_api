package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/todo/dashboard"
)

type DashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(conn *PostgreSQLConnector) (repo DashboardRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = DashboardRepository{
		db: db,
	}
	return
}

func (d DashboardRepository) GetDashboardSummary(userID int) (summary dashboard.DashboardSummary, err error) {
	table := "task"

	query := fmt.Sprintf(`
	select
		coalesce(count(t.id), 0) as tasks_count,
		coalesce(sum(case when is_done = true then 1 else 0 end), 0) as done_tasks_count,
		min(a.name) as name,
		min(a.last_name) as last_name,
		min(a.nickname) as nickname
	from
		%s t
	inner join account a on t.created_by = a.id
	where
		t.created_by = $1
	group by t.created_by, a.name, a.last_name, a.nickname
	`, table)

	err = d.db.QueryRow(query, userID).Scan(
		&summary.TasksCount, &summary.DoneTasksCount,
		&summary.Name, &summary.LastName, &summary.Nickname,
	)
	if err != nil {
		err = errorInRow(table, "get", err)
	}
	return
}
