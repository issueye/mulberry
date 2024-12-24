package service

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	commonModel "carambola/host/common/model"
	"carambola/host/common/service"

	"gorm.io/gorm"
)

type Rule struct {
	service.BaseService[model.RuleInfo]
}

func NewRule(args ...any) *Rule {
	srv := &Rule{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListRule
// 根据条件查询列表
func (r *Rule) ListRule(condition *commonModel.PageQuery[*requests.QueryRule]) (*commonModel.ResPage[model.RuleInfo], error) {
	return service.GetList[model.RuleInfo](condition, func(qu *requests.QueryRule, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (r *Rule) SaveCode(id uint, code string) error {
	return r.UpdateByMap(id, map[string]any{
		"script_content": code,
	})
}

func (r *Rule) UpdateRuleStatus(id uint, status int) error {
	return r.UpdateByMap(id, map[string]any{
		"status": status,
	})
}
