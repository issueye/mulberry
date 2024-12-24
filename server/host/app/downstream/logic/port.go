package logic

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	"carambola/host/app/downstream/service"
	commonModel "carambola/host/common/model"
	"carambola/host/global"
)

func CreatePort(req *requests.CreatePort) error {
	srv := service.NewPort(global.DB, false)
	return srv.Create(&req.PortInfo)
}

func UpdatePort(req *requests.UpdatePort) error {
	return service.NewPort(global.DB, false).Update(req.ID, &req.PortInfo)
}

func DeletePort(id uint) error {
	return service.NewPort(global.DB, false).Delete(id)
}

func PortList(condition *commonModel.PageQuery[*requests.QueryPort]) (*commonModel.ResPage[model.PortInfo], error) {
	return service.NewPort(global.DB, false).ListPort(condition)
}

func GetPort(id uint) (*model.PortInfo, error) {
	return service.NewPort(global.DB, false).GetById(id)
}

func SaveCode(id uint, code string) error {
	return service.NewPort(global.DB, false).SaveCode(id, code)
}
