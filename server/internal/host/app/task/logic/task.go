package logic

import (
	"fmt"
	"mulberry/internal/global"
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/task/model"
	"mulberry/internal/host/app/task/requests"
	"mulberry/internal/host/app/task/service"
	"mulberry/internal/host/ipc"
	"mulberry/pkg/ipc/grpc/pb"
	"strconv"
)

func CreateTask(req *requests.CreateTask) error {
	srv := service.NewTask(global.DB, false)
	return srv.Create(&req.Task)
}

func UpdateTask(req *requests.UpdateTask) error {
	clientInfo, err := service.NewClient(global.DB, false).GetByField("client_auth_id", req.ClientAuthId)
	if err != nil {
		return err
	}

	if clientInfo.Status == 1 {
		global.Logger.Sugar().Errorf("客户端正在运行，不能修改任务", clientInfo.Name)
	}

	return service.NewTask(global.DB, false).Update(req.ID, &req.Task)
}

func DeleteTask(id uint) error {
	return service.NewTask(global.DB, false).Delete(id)
}

func TaskList(condition *commonModel.PageQuery[*requests.QueryTask]) (*commonModel.ResPage[model.Task], error) {
	return service.NewTask(global.DB, false).ListTask(condition)
}

func GetTask(id uint) (*model.Task, error) {
	return service.NewTask(global.DB, false).GetById(id)
}

func SaveCode(id uint, code string) error {
	return service.NewTask(global.DB, false).SaveCode(id, code)
}

func UpdateTaskStatus(id uint) error {
	status := 0
	srv := service.NewTask(global.DB, false)
	task, err := srv.GetById(id)
	if err != nil {
		return err
	}

	// 获取客户端信息
	client := service.NewClient(global.DB, false)
	clientInfo, err := client.GetByField("client_auth_id", task.ClientAuthId)
	if err != nil {
		return err
	}

	if clientInfo.Status == 0 {
		global.Logger.Sugar().Errorf("客户端: %s 未开启", clientInfo.Name)
		return fmt.Errorf("客户端: %s 未开启", clientInfo.Name)
	}

	helper := ipc.GetCommonHelper(task.ClientAuthId)

	info := &pb.TaskInfo{
		Id:             strconv.FormatInt(int64(task.ID), 10),
		Name:           task.Name,
		Description:    task.Description,
		CronExpression: task.CronExpression,
		ScriptLanguage: string(*task.ScriptLanguage),
		ScriptContent:  task.ScriptContent,
		Status:         int32(status),
	}

	if task.Status == 0 {
		status = 1
		info.Status = 1
		global.Logger.Sugar().Debugf("启动定时任务: %s", info.Name)
		helper.CommandStream.Send(&pb.CommandResponse{Command: pb.CommandType_Start, Task: info})
	}

	if task.Status == 1 {
		status = 0
		info.Status = 0
		global.Logger.Sugar().Debugf("关闭定时任务: %s", info.Name)
		helper.CommandStream.Send(&pb.CommandResponse{Command: pb.CommandType_Stop, Task: info})
	}

	return srv.UpdateTaskStatus(id, status)
}

func RunTask(id uint) error {
	srv := service.NewTask(global.DB, false)
	task, err := srv.GetById(id)
	if err != nil {
		return err
	}

	// 获取客户端信息
	client := service.NewClient(global.DB, false)
	clientInfo, err := client.GetByField("client_auth_id", task.ClientAuthId)
	if err != nil {
		return err
	}

	if clientInfo.Status == 0 {
		global.Logger.Sugar().Errorf("客户端: %s 未开启", clientInfo.Name)
		return fmt.Errorf("客户端: %s 未开启", clientInfo.Name)
	}

	info := &pb.TaskInfo{
		Id:             strconv.FormatInt(int64(task.ID), 10),
		Name:           task.Name,
		Description:    task.Description,
		CronExpression: task.CronExpression,
		ScriptLanguage: string(*task.ScriptLanguage),
		ScriptContent:  task.ScriptContent,
		Status:         int32(task.Status),
	}

	global.Logger.Sugar().Debugf("运行定时任务: %s", info.Name)

	helper := ipc.GetCommonHelper(task.ClientAuthId)
	if helper.CommandStream == nil {
		return fmt.Errorf("[%s]客户端: %s 未连接", clientInfo.ClientAuthId, clientInfo.Name)
	}

	return helper.CommandStream.Send(&pb.CommandResponse{Command: pb.CommandType_Run, Task: info})
}
