package dashboard

import "github.com/coffemanfp/todo/account"

type DashboardSummary struct {
	account.AccountBasic
	TasksCount     int
	DoneTasksCount int
}
