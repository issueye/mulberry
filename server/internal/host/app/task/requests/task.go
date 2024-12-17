package requests

import (
	"encoding/json"
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/task/model"
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
	KeyWords string `json:"keywords" form:"keywords"`   // 关键词
	IsHidden int    `json:"is_hidden" form:"is_hidden"` // 0 不隐藏 1 隐藏
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
