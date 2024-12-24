package logic

import (
	"carambola/host/app/downstream/model"
	"carambola/host/app/downstream/requests"
	"carambola/host/app/downstream/service"
	commonModel "carambola/host/common/model"
	"carambola/host/global"
)

func CreatePage(req *requests.CreatePage) error {
	srv := service.NewPage(global.DB, false)
	return srv.Create(&req.PageInfo)
}

func UpdatePage(req *requests.UpdatePage) error {
	return service.NewPage(global.DB, false).Update(req.ID, &req.PageInfo)
}

func DeletePage(id uint) error {
	return service.NewPage(global.DB, false).Delete(id)
}

func PageList(condition *commonModel.PageQuery[*requests.QueryPage]) (*commonModel.ResPage[model.PageInfo], error) {
	return service.NewPage(global.DB, false).ListPage(condition)
}

func GetPage(id uint) (*model.PageInfo, error) {
	return service.NewPage(global.DB, false).GetById(id)
}
