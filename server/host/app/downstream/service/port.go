package service

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"

	"gorm.io/gorm"
)

type Port struct {
	service.BaseService[model.PortInfo]
}

func NewPort(args ...any) *Port {
	srv := &Port{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListPort
// 根据条件查询列表
func (r *Port) ListPort(condition *commonModel.PageQuery[*requests.QueryPort]) (*commonModel.ResPage[model.PortInfo], error) {
	return service.GetList[model.PortInfo](condition, func(qu *requests.QueryPort, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *Port) SaveCode(id uint, code string) error {
	return r.UpdateByMap(id, map[string]any{
		"script_content": code,
	})
}

func (r *Port) UpdatePortStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
