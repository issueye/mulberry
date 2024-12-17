package logic

import (
	"errors"
	"mulberry/internal/global"
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/task/model"
	"mulberry/internal/host/app/task/requests"
	"mulberry/internal/host/app/task/service"
)

func CreateDatabaseInfo(req *requests.CreateDatabase) error {
	srv := service.NewDatabaseInfo(global.DB, false)
	info, err := srv.GetByField("name", req.Name)
	if err != nil {
		return err
	}

	if info.ID > 0 {
		return errors.New("名称已经存在")
	}

	err = srv.Create(&req.DatabaseInfo)
	if err != nil {
		return err
	}

	_, err = global.RedkaDB.List().PushBack(global.DB_Key, req.DatabaseInfo.ToJson())
	return err
}

func UpdateDatabaseInfo(req *requests.UpdateDatabaseInfo) error {
	err := service.NewDatabaseInfo(global.DB, false).Update(req.ID, &req.DatabaseInfo)
	if err != nil {
		return err
	}

	index, err := GetIndexDB(req.ID)
	if err != nil {
		return err
	}

	if index == -1 {
		return nil
	}

	global.RedkaDB.List().Set(global.DB_Key, index, req.DatabaseInfo.ToJson())
	return nil
}

func GetIndexDB(id uint) (int, error) {
	list, err := global.RedkaDB.List().Range(global.DB_Key, 0, -1)
	if err != nil {
		return -1, err
	}

	for index, value := range list {
		db := model.DatabaseInfo{}.FromJson(value.String())
		if db != nil {
			if db.ID == id {
				return index, nil
			}
		}
	}

	return -1, nil
}

func DeleteDatabaseInfo(id uint) error {
	srv := service.NewDatabaseInfo(global.DB, false)
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	if info.ID == 0 {
		return errors.New("未找到数据")
	}

	err = service.NewDatabaseInfo(global.DB, false).Delete(id)
	if err != nil {
		return err
	}

	global.RedkaDB.List().Delete(global.DB_Key, info.ToJson())
	return nil
}

func DatabaseInfoList(condition *commonModel.PageQuery[*requests.QueryDatabaseInfo]) (*commonModel.ResPage[model.DatabaseInfo], error) {
	return service.NewDatabaseInfo(global.DB, false).ListDatabaseInfo(condition)
}

func GetDatabaseInfo(id uint) (*model.DatabaseInfo, error) {
	info, err := service.NewDatabaseInfo(global.DB, false).GetById(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func InitTaskDB() {
	// 将所有客户端状态修改为未启动
	if err := service.NewClient(global.DB, false).UpdateByField("1", 1, map[string]any{"status": "0"}); err != nil {
		return
	}

	// 将所有任务状态修改为未开启
	if err := service.NewTask(global.DB, false).UpdateByField("1", 1, map[string]any{"status": "0"}); err != nil {
		return
	}

	// 先清理数据
	dbList, err := global.RedkaDB.List().Range(global.DB_Key, 0, -1)
	if err != nil {
		return
	}

	for _, db := range dbList {
		global.RedkaDB.List().Delete(global.DB_Key, db.String())
	}

	// 获取数据库列表
	list, err := service.NewDatabaseInfo(global.DB, false).GetAll()
	if err != nil {
		return
	}

	for _, value := range list {
		_, err = global.RedkaDB.List().PushBack(global.DB_Key, value.ToJson())
		if err != nil {
			return
		}
	}
}
