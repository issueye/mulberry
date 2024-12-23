package requests

import (
	"carambola/host/app/task/model"
	commonModel "carambola/host/common/model"
)

type CreateDatabase struct {
	model.DatabaseInfo
}

func NewCreateDatabase() *CreateDatabase {
	return &CreateDatabase{
		DatabaseInfo: model.DatabaseInfo{},
	}
}

type UpdateDatabaseInfo struct {
	model.DatabaseInfo
}

func NewUpdateDatabaseInfo() *UpdateDatabaseInfo {
	return &UpdateDatabaseInfo{
		DatabaseInfo: model.DatabaseInfo{},
	}
}

type QueryDatabaseInfo struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryDatabaseInfo() *commonModel.PageQuery[*QueryDatabaseInfo] {
	return commonModel.NewPageQuery(0, 0, &QueryDatabaseInfo{})
}
