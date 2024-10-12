package database

import "github.com/coffemanfp/todo/dashboard"

// DASHBOARD_REPOSITORY is the key to be used when creating the repositories hashmap.
const DASHBOARD_REPOSITORY RepositoryID = "DASHBOARD"

// DashboardRepository defines the behaviors to be used by a DashboardRepository implementation.
type DashboardRepository interface {
	GetDashboardSummary(userID int) (summary dashboard.DashboardSummary, err error)
}
