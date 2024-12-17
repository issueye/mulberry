package service

import (
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/common/service"
	"mulberry/internal/host/app/task/model"
	"mulberry/internal/host/app/task/requests"

	"gorm.io/gorm"
)

type Client struct {
	service.BaseService[model.Client]
}

func NewClient(args ...any) *Client {
	srv := &Client{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListClient
// 根据条件查询列表
func (r *Client) ListClient(condition *commonModel.PageQuery[*requests.QueryClient]) (*commonModel.ResPage[model.Client], error) {
	return service.GetList[model.Client](condition, func(qu *requests.QueryClient, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}
