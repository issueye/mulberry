package service

import (
	"carambola/host/app/task/model"
	"carambola/host/common/service"
)

type History struct {
	service.BaseService[model.TaskExecutionHistory]
}

func NewHistory(args ...any) *History {
	srv := &History{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}
