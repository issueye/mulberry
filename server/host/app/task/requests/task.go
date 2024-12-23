package requests

import (
	"carambola/host/app/task/model"
	commonModel "carambola/host/common/model"
	"encoding/json"
)

type CreateTask struct {
	model.Task
}

func NewCreateTask() *CreateTask {
	return &CreateTask{
		Task: model.Task{},
	}
}

type UpdateTask struct {
	model.Task
}

func NewUpdateTask() *UpdateTask {
	return &UpdateTask{
		Task: model.Task{},
	}
}

func (req *CreateTask) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryTask struct {
	KeyWords     string `json:"keywords" form:"keywords"`             // 关键词
	IsHidden     int    `json:"is_hidden" form:"is_hidden"`           // 0 不隐藏 1 隐藏
	ClientAuthId string `json:"client_auth_id" form:"client_auth_id"` // 客户端认证ID
}

func NewQueryTask() *commonModel.PageQuery[*QueryTask] {
	return commonModel.NewPageQuery(0, 0, &QueryTask{})
}

type SaveCode struct {
	ID   uint   `json:"id" form:"id"`
	Code string `json:"code" form:"code"`
}

func NewSaveCode() *SaveCode {
	return &SaveCode{}
}

type QueryTaskLog struct {
	TaskID   int `json:"task_id" form:"task_id"`     // 任务id
	ClientID int `json:"client_id" form:"client_id"` // 客户端id
}

func NewQueryTaskLog() *commonModel.PageQuery[*QueryTaskLog] {
	return commonModel.NewPageQuery(0, 0, &QueryTaskLog{})
}
