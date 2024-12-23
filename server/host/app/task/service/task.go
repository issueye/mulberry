package service

import (
	"carambola/host/app/task/model"
	"carambola/host/app/task/requests"
	"carambola/host/app/task/response"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"
	"fmt"

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

		if qu.ClientAuthId != "" {
			d = d.Where("client_auth_id =?", qu.ClientAuthId)
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

func (r *Task) GetTaskLogList(req *commonModel.PageQuery[*requests.QueryTaskLog]) (*commonModel.ResPage[response.TaskLogResponse], error) {
	var sql_str = `select 
h.*,
info.name as task_name,
client.id as c_id,
client.name as client_name 
from task_execution_history h 
left join task_client_info client on h.client_id = client.client_auth_id 
left join task_info info on h.task_id = info.id`
	list := make([]*response.TaskLogResponse, 0)

	qry := r.DB.Table(fmt.Sprintf("(%s)tb", sql_str))
	// sql := qry.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(nil) })
	// fmt.Println("sql", sql)

	qry = qry.Order("end_time desc")

	if req.Condition.ClientID > 0 {
		qry = qry.Where("c_id = ?", req.Condition.ClientID)
	}

	if req.Condition.TaskID > 0 {
		qry = qry.Where("task_id = ?", req.Condition.TaskID)
	}

	total := int64(0)
	err := qry.Count(&total).Error
	if err != nil {
		return nil, err
	}

	if req.PageNum == 0 || req.PageSize == 0 {
		err = qry.Find(&list).Error
	} else {
		err = qry.
			Limit(req.PageSize).
			Offset((req.PageNum - 1) * req.PageSize).
			Find(&list).Error
	}

	resData := commonModel.NewResPage(req.PageNum, req.PageSize, int(total), list)
	return resData, err
}
