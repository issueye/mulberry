package service

import (
	"carambola/host/app/task/model"
	"carambola/host/app/task/requests"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"

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
