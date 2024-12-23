package service

import (
	"carambola/host/app/task/model"
	"carambola/host/app/task/requests"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"

	"gorm.io/gorm"
)

type DatabaseInfo struct {
	service.BaseService[model.DatabaseInfo]
}

func NewDatabaseInfo(args ...any) *DatabaseInfo {
	srv := &DatabaseInfo{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListDatabaseInfo
// 根据条件查询列表
func (r *DatabaseInfo) ListDatabaseInfo(condition *commonModel.PageQuery[*requests.QueryDatabaseInfo]) (*commonModel.ResPage[model.DatabaseInfo], error) {
	return service.GetList[model.DatabaseInfo](condition, func(qu *requests.QueryDatabaseInfo, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}
