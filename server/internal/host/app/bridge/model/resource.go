package model

import "mulberry/internal/host/app/common/model"

// ResourceInfo
// 资源信息
type ResourceInfo struct {
	model.BaseModel
	ResourceBase
}

type ResourceBase struct {
	Title    string `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"` // 标题
	FileName string `label:"文件名称" gorm:"column:file_name;size:2000;comment:文件名称;" json:"fileName"`        // 文件名称
	Ext      string `label:"文件类型" gorm:"column:ext;size:200;comment:文件类型;" json:"ext"`                    // 文件类型
	Mark     string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                     // 备注
}

func (mod *ResourceInfo) Copy(data *ResourceBase) {
	mod.Title = data.Title
	mod.FileName = data.FileName
	mod.Ext = data.Ext
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (ResourceInfo) TableName() string { return "bridge_resource_info" }
