package route

import (
	v1 "mulberry/internal/host/app/task/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	task := r.Group("task")
	{
		task.POST("", v1.CreateTask)
		task.PUT("", v1.UpdateTask)
		task.DELETE(":id", v1.DeleteTask)
		task.POST("list", v1.TaskList)
		task.GET(":id", v1.GetTask)
		task.PUT("save_code", v1.SaveCode)
		task.PUT("updateStatus/:id", v1.UpdateTaskStatus)
		task.GET("run/:id", v1.RunTask)
	}

	client := r.Group("client")
	{
		client.POST("", v1.CreateClient)
		client.PUT("", v1.UpdateClient)
		client.DELETE(":id", v1.DeleteClient)
		client.POST("list", v1.ClientList)
		client.GET(":id", v1.GetClient)
		client.PUT("close/:id", v1.CloseClient)
		client.PUT("open/:id", v1.OpenClient)
	}

	database := r.Group("db")
	{
		database.POST("", v1.CreateDatabaseInfo)
		database.PUT("", v1.UpdateDatabaseInfo)
		database.DELETE(":id", v1.DeleteDatabaseInfo)
		database.POST("list", v1.DatabaseInfoList)
		database.GET(":id", v1.GetDatabaseInfo)
	}
}
