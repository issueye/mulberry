package route

import (
	adminRoute "carambola/host/app/admin/route"
	taskRoute "carambola/host/app/task/route"
	"carambola/host/common/controller"
	"carambola/host/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctl := controller.New(ctx)
			ctl.SuccessData(map[string]any{"msg": "pong"})
		})

		// 注册管理路由
		adminRoute.Register(v1)
		// 注册业务路由
		taskRoute.Register(v1)
	}

	r.NoRoute(func(ctx *gin.Context) {
		global.Logger.Logger.Error("404", zap.String("path", ctx.Request.URL.Path), zap.String("method", ctx.Request.Method))
		ctl := controller.New(ctx)
		ctl.FailWithCode(http.StatusNotFound, "not found")
	})
}