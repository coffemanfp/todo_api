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

	accountRepo, err := psql.NewAccountRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	db.Repositories = map[database.RepositoryID]interface{}{
		database.ACCOUNT_REPOSITORY: accountRepo,
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
