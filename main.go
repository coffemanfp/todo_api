package main

import (
	"fmt"
	"log"

	"github.com/coffemanfp/todo/config"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/database/psql"
	"github.com/coffemanfp/todo/server"
	"github.com/coffemanfp/todo/server/gin"
)

var (
	conf         config.ConfigInfo
	serverEngine server.Engine
)

func main() {
	serverEngine.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func setUpDatabase(conf config.ConfigInfo) (db database.Database, err error) {
	db.Conn = psql.NewPostgreSQLConnector(
		conf.PostgreSQLProperties.User,
		conf.PostgreSQLProperties.Password,
		conf.PostgreSQLProperties.Name,
		conf.PostgreSQLProperties.Host,
		conf.PostgreSQLProperties.Port,
	)

	err = db.Conn.Connect()
	if err != nil {
		return
	}

	authRepo, err := psql.NewAuthRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	listRepo, err := psql.NewListRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	taskRepo, err := psql.NewTaskRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}
	categoryRepo, err := psql.NewCategoryRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}
	dashboardRepo, err := psql.NewDashboardRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	db.Repositories = map[database.RepositoryID]interface{}{
		database.AUTH_REPOSITORY:      authRepo,
		database.LIST_REPOSITORY:      listRepo,
		database.TASK_REPOSITORY:      taskRepo,
		database.CATEGORY_REPOSITORY:  categoryRepo,
		database.DASHBOARD_REPOSITORY: dashboardRepo,
	}
	return
}

func init() {
	var err error
	conf, err = config.NewEnvManagerConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := setUpDatabase(conf)
	if err != nil {
		log.Fatal(err)
	}

	serverEngine = gin.New(conf, db)
}
