package routes

import (
	"main/ToDoList/api"
	"main/ToDoList/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//增
			authed.POST("task", api.CreateTask)

			// 查
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks", api.ListTask)
			authed.GET("tasks/finish", api.FinishTask)
			authed.GET("tasks/unfinish", api.UnfinishTask)
			authed.POST("search", api.SearchTask)
			// 改
			authed.PUT("task/:id", api.UpdateTask)
			authed.PUT("task/change", api.ChangeTask)
			// 删
			authed.DELETE("task/:id", api.DeleteTask)
			authed.DELETE("task/finish", api.DeleteFinishTask)
			authed.DELETE("task/unfinish", api.DeleteUnfinishTask)
			authed.DELETE("task/all", api.DeleteAllTask)
		}
	}
	return r
}
