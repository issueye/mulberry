package initialize

import (
	"carambola/client/global"
	"carambola/client/ipc"
	"carambola/client/task"
	"runtime/debug"
)

func Init(clientID string) {
	global.ClientID = clientID
	// 初始化IPC客户端
	ipc.InitIpc()
	// 初始化配置
	InitRuntime()
	// 初始化日志
	InitLogger()

	defer func() {
		err := recover()
		if err != nil {
			// 获取堆栈信息
			debug.PrintStack()
			global.Logger.Sugar().Errorf("初始化失败：%s", err)
			return
		}
	}()

	// 初始化 task
	task.InitTask()
	// 初始化脚本引擎
	InitCodeEngine()
	// 初始化数据
	InitDB()
	// 启动任务
	ipc.CommandTask()
	select {}
}
