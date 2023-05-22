package handlers

import (
	"github.com/coffemanfp/todo/config"
	"github.com/coffemanfp/todo/database"
)

var db database.Repositories
var conf config.ConfigInfo
var handlers = map[string]Handler{
	"login": Login{},
}

func Init(newDb database.Repositories, newConf config.ConfigInfo) {
	db = newDb
	conf = newConf
}
