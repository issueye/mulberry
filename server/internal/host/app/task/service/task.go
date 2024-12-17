package service

import (
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/common/service"
	"mulberry/internal/host/app/task/model"
	"mulberry/internal/host/app/task/requests"

	"gorm.io/gorm"
)

type Task struct {
	service.BaseService[model.Task]
}

func NewTask(args ...any) *Task {
	srv := &Task{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListTask
// 根据条件查询列表
func (r *Task) ListTask(condition *commonModel.PageQuery[*requests.QueryTask]) (*commonModel.ResPage[model.Task], error) {
	return service.GetList[model.Task](condition, func(qu *requests.QueryTask, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *Task) SaveCode(id uint, code string) error {
	return r.UpdateByMap(id, map[string]any{
		"script_content": code,
	})
}

func (r *Task) UpdateTaskStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
