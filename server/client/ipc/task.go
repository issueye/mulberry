package ipc

import (
	"carambola/client/global"
	"carambola/client/ipc/register"
	"carambola/client/task"
	"carambola/common/ipc/grpc/pb"
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
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
		var err error

		guid := uuid.New().String()
		global.Logger.Sugar().Infof("开始执行任务：%s 执行ID: %s", info.Name, guid)

		// 发送任务开始消息
		global.IpcClient.HostHelper().TaskStart(context.Background(), &pb.TaskCallback{
			ClientID:    global.ClientID,
			Id:          info.Id,
			Name:        info.Name,
			ExecutionID: guid,
		})

		defer func() {
			code := pb.RES_CODE_SUCCESS
			msg := "任务执行成功"
			if err != nil {
				code = pb.RES_CODE_FAILED
				msg = err.Error()
			}

			// 获取下一次执行时间
			entry := task.GetTaskCron().GetEntry(info.Id)
			// 发送任务完成消息
			global.IpcClient.HostHelper().TaskComplete(context.Background(), &pb.TaskCallback{
				ClientID:          global.ClientID,
				Id:                info.Id,
				Name:              info.Name,
				ExecutionID:       guid,
				NextExecutionTime: entry.Next.Format("2006-01-02 15:04:05"),
				Code:              code,
				Message:           msg,
			})
			global.Logger.Sugar().Infof("结束执行任务：%s 执行ID: %s", info.Name, guid)
		}()

		vm := global.CodeEngine.GetRuntime()
		defer global.CodeEngine.PutRuntime(vm)

		// pdf 文件转字符串
		vm.Set("fileToBase64", register.FileToBase64)
		vm.Set("runApp", register.RunApp)

		err = vm.RunCode(info.ScriptContent)
		if err != nil {
			global.Logger.Sugar().Errorf("运行脚本失败：%s", err.Error())
			return
		}
	}
}
