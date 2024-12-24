package logic

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	"carambola/host/app/downstream/service"
	commonModel "carambola/host/common/model"
	"carambola/host/global"
)

func CreateRule(req *requests.CreateRule) error {
	srv := service.NewRule(global.DB, false)
	return srv.Create(&req.RuleInfo)
}

func UpdateRule(req *requests.UpdateRule) error {
	return service.NewRule(global.DB, false).Update(req.ID, &req.RuleInfo)
}

func DeleteRule(id uint) error {
	return service.NewRule(global.DB, false).Delete(id)
}

func RuleList(condition *commonModel.PageQuery[*requests.QueryRule]) (*commonModel.ResPage[model.RuleInfo], error) {
	return service.NewRule(global.DB, false).ListRule(condition)
}

func GetRule(id uint) (*model.RuleInfo, error) {
	return service.NewRule(global.DB, false).GetById(id)
}
