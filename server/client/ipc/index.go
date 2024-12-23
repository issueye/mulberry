package ipc

import (
	"carambola/client/global"
	"carambola/common/ipc/client"
	"carambola/common/ipc/vars"
)

func InitIpc() {
	vars.PIPE_NAME = vars.PIPE_NAME + "_carambola"

	c, err := client.NewClient()
	if err != nil {
		panic(err)
	}

	global.IpcClient = c
	global.IpcClient.ClientID = global.ClientID

	vars.AppName = "任务客户端"
	vars.ClientID = global.ClientID
	vars.Version = "1.0.0.1"
	global.IpcClient.Register()
	go global.IpcClient.Heartbeat()
}
