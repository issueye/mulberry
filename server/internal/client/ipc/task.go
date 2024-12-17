package ipc

import (
	"context"
	"errors"
	"io"
	"mulberry/internal/client/ipc/register"
	"mulberry/internal/client/task"
	"mulberry/internal/global"
	"mulberry/pkg/ipc/grpc/pb"
)

func CommandTask() {
	global.Logger.Sugar().Infof("开始接收命令")

	commandStream, err := global.IpcClient.HostHelper().Command(context.Background(), &pb.ClientRequest{ClientID: global.ClientID})
	if err != nil {
		global.Logger.Sugar().Errorf("连接失败：%s", err.Error())
		return
	}

	for {
		command, err := commandStream.Recv()
		if errors.Is(err, io.EOF) {
			global.Logger.Sugar().Errorf("接收命令失败：%s", err.Error())
			break
		}

		if err != nil {
			global.Logger.Sugar().Errorf("接收命令失败：%s", err.Error())
			continue
		}

		global.Logger.Sugar().Infof("接收到命令：%s", command.Command.String())

		switch command.Command {
		case pb.CommandType_Run:
			TaskRun(command.Task)
		case pb.CommandType_Stop:
			TaskStop(command.Task)
		case pb.CommandType_Start:
			TaskStart(command.Task)
		}
	}
}

func TaskRun(info *pb.TaskInfo) {
	task.GetTaskCron().Run(info.Name, taskFunc(info))
}

func TaskStop(info *pb.TaskInfo) {
	task.GetTaskCron().Remove(info.Id)
}

func TaskStart(info *pb.TaskInfo) {
	task.GetTaskCron().AddFuncAt(info.Id, info.CronExpression, taskFunc(info))
}

func taskFunc(info *pb.TaskInfo) func() {
	return func() {
		vm := global.CodeEngine.GetRuntime()
		defer global.CodeEngine.PutRuntime(vm)

		// pdf 文件转字符串
		vm.Set("fileToBase64", register.FileToBase64)
		vm.Set("runApp", register.RunApp)

		err := vm.RunCode(info.ScriptContent)
		if err != nil {
			global.Logger.Sugar().Errorf("运行脚本失败：%s", err.Error())
			return
		}
	}
}
