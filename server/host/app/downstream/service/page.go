package service

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"

	"gorm.io/gorm"
)

type Page struct {
	service.BaseService[model.PageInfo]
}

func NewPage(args ...any) *Page {
	srv := &Page{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListPage
// 根据条件查询列表
func (r *Page) ListPage(condition *commonModel.PageQuery[*requests.QueryPage]) (*commonModel.ResPage[model.PageInfo], error) {
	return service.GetList[model.PageInfo](condition, func(qu *requests.QueryPage, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *Page) SaveCode(id uint, code string) error {
	return r.UpdateByMap(id, map[string]any{
		"script_content": code,
	})
}

func (r *Page) UpdatePageStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
