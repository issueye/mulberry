package main

import (
	"carambola/host/global"
	"carambola/host/initialize"
	"carambola/host/pages/home"
	_ "carambola/host/winappres"
	"embed"
	"io/fs"

	"github.com/ying32/govcl/vcl"
)

//	@title			定时任务调度服务系统v1.0
//	@version		V1.1
//	@description	golang 定时任务调度服务系统v1.0

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//go:embed web/*
var webDir embed.FS

func main() {
	staticFp, _ := fs.Sub(webDir, "web")
	global.StaticWEB = staticFp
	initialize.InitRuntime()
	vcl.RunApp(&home.FrmHome)
}
