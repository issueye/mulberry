package route

import (
	v1 "carambola/host/app/downstream/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	port := r.Group("port")
	{
		port.POST("", v1.CreatePort)
		port.PUT("", v1.UpdatePort)
		port.DELETE(":id", v1.DeletePort)
		port.POST("list", v1.PortList)
		port.GET(":id", v1.GetPort)
	}

	rule := r.Group("rule")
	{
		rule.POST("", v1.CreateRule)
		rule.PUT("", v1.UpdateRule)
		rule.DELETE(":id", v1.DeleteRule)
		rule.POST("list", v1.RuleList)
		rule.GET(":id", v1.GetRule)
	}

	page := r.Group("page")
	{
		page.POST("", v1.CreatePage)
		page.PUT("", v1.UpdatePage)
		page.DELETE(":id", v1.DeletePage)
		page.POST("list", v1.PageList)
		page.GET(":id", v1.GetPage)
	}
}
