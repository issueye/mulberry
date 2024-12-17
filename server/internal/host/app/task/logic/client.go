package logic

import (
	"errors"
	"mulberry/internal/global"
	commonModel "mulberry/internal/host/app/common/model"
	"mulberry/internal/host/app/task/model"
	"mulberry/internal/host/app/task/requests"
	"mulberry/internal/host/app/task/service"
)

func CreateClient(req *requests.CreateClient) error {
	srv := service.NewClient(global.DB, false)
	info, err := srv.GetByField("client_auth_id", req.Client.ClientAuthId)
	if err != nil {
		return err
	}

	if info.ID > 0 {
		return errors.New("客户端已存在")
	}

	return srv.Create(&req.Client)
}

func UpdateClient(req *requests.UpdateClient) error {
	return service.NewClient(global.DB, false).Update(req.ID, &req.Client)
}

func DeleteClient(id uint) error {
	return service.NewClient(global.DB, false).Delete(id)
}

func ClientList(condition *commonModel.PageQuery[*requests.QueryClient]) (*commonModel.ResPage[model.Client], error) {
	return service.NewClient(global.DB, false).ListClient(condition)
}

func GetClient(id uint) (*model.Client, error) {
	return service.NewClient(global.DB, false).GetById(id)
}

func CloseClient(id uint) error {
	tx := global.DB.Begin()

	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	srv := service.NewClient(tx, true)

	var info *model.Client
	info, err = srv.GetById(id)
	if err != nil {
		return err
	}

	// 判断是否正在运行
	if info.Status != 1 {
		return nil
	}

	err = srv.UpdateByMap(id, map[string]any{
		"status": 0,
	})

	if err != nil {
		return err
	}

	err = CloseProcess(info.ClientAuthId)
	return err
}

func RunClient(id uint) error {
	srv := service.NewClient(global.DB, false)
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	// 判断是否正在运行
	if info.Status == 1 {
		return nil
	}

	err = service.NewClient(global.DB, false).UpdateByMap(id, map[string]any{
		"status": 1,
	})

	if err != nil {
		return err
	}

	return RunProcess(info.ClientAuthId)
}
