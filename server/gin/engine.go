package gin

import (
	"github.com/coffemanfp/todo/config"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server"
	"github.com/coffemanfp/todo/server/gin/handlers"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	conf config.ConfigInfo
	db   database.Database
	r    *gin.Engine
}

func New(conf config.ConfigInfo, db database.Database) server.Engine {
	ge := GinEngine{
		conf: conf,
		db:   db,
		r:    gin.New(),
	}

	handlers.Init(ge.db.Repositories, ge.conf)

	ge.r.Use(newCors(ge.conf))
	ge.r.Use(errorHandler())
	v1 := ge.r.Group("/v1")

	ge.setCommonMiddlewares(v1)
	ge.setAuthHandlers(v1)
	ge.setListHandlers(v1)
	ge.setTaskHandlers(v1)
	ge.setSearchHandlers(v1)
	ge.setCategoriesHandlers(v1)
	ge.setDashboardHandlers(v1)
	return ge.r
}

func (ge GinEngine) setAuthHandlers(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", handlers.Login{}.Do)
	auth.POST("/register", handlers.Register{}.Do)
}

func (ge GinEngine) setListHandlers(r *gin.RouterGroup) {
	list := r.Group("/lists")
	list.Use(authorize(ge.conf.Server.SecretKey))
	list.GET("/:id", handlers.GetList{}.Do)
	list.GET("", handlers.GetSomeLists{}.Do)
	list.POST("", handlers.CreateList{}.Do)
	list.PUT("/:id", handlers.UpdateList{}.Do)
	list.DELETE("/:id", handlers.DeleteList{}.Do)
}

func (ge GinEngine) setTaskHandlers(r *gin.RouterGroup) {
	task := r.Group("/tasks")
	task.Use(authorize(ge.conf.Server.SecretKey))
	task.GET("", handlers.GetSomeTasks{}.Do)
	task.GET("/:id", handlers.GetTask{}.Do)
	task.POST("", handlers.CreateTask{}.Do)
	task.PUT("/:id", handlers.UpdateTask{}.Do)
	task.DELETE("/:id", handlers.DeleteTask{}.Do)
}

func (ge GinEngine) setCategoriesHandlers(r *gin.RouterGroup) {
	category := r.Group("/categories")
	category.Use(authorize(ge.conf.Server.SecretKey))
	category.GET("", handlers.GetSomeCategories{}.Do)
	category.POST("", handlers.CreateCategory{}.Do)
	category.PUT("/:id", handlers.UpdateCategory{}.Do)
	category.DELETE("/:id", handlers.DeleteCategory{}.Do)
	category.POST("/binds", handlers.CreateCategoryBind{}.Do)
}

// setSearchHandlers configures search-related routes and handlers.
func (ge GinEngine) setSearchHandlers(r *gin.RouterGroup) {
	// Create a sub-group for search routes
	product := r.Group("/search")
	// Use authorization middleware to protect this route
	product.Use(authorize(ge.conf.Server.SecretKey))
	// Configure endpoint for searching products
	product.GET("", handlers.Search{}.Do)
}

func (ge GinEngine) setDashboardHandlers(r *gin.RouterGroup) {
	// Create a sub-group for dashboard routes
	dashboard := r.Group("/users/:id/dashboard")
	// Use authorization middleware to protect this route
	dashboard.Use(authorize(ge.conf.Server.SecretKey))
	// Configure endpoint for getting dashboard summary
	dashboard.GET("", handlers.GetDashboardSummary{}.Do)
}

func (ge GinEngine) setCommonMiddlewares(r *gin.RouterGroup) {
	r.Use(gin.Recovery())
	r.Use(logger())
}
