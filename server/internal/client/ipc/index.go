package ipc

import (
	"mulberry/internal/global"
	"mulberry/pkg/ipc/client"
	"mulberry/pkg/ipc/vars"
)

func InitIpc() {
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
