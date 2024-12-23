package model

import "carambola/host/common/model"

type PageInfo struct {
	model.BaseModel
	PageBase
}

type PageBase struct {
	Name            string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`                                   // 名称
	Title           string `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"`                                 // 标题
	Version         string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`                              // 版本
	PortId          string `binding:"required" label:"端口号" gorm:"column:port_id;type:int;comment:端口信息编码;" json:"port_id"`                        // 端口信息编码
	ProductCode     string `binding:"required" label:"产品代码" gorm:"column:product_code;size:200;comment:产品代码;" json:"product_code"`               // 产品代码
	Thumbnail       string `binding:"required" label:"缩略图" gorm:"column:thumbnail;size:200;comment:缩略图;" json:"thumbnail"`                       // 缩略图
	UseVersionRoute int    `binding:"required" label:"使用版本路由" gorm:"column:use_version_route;type:int;comment:使用版本路由;" json:"use_version_route"` // 使用版本路由
	Status          bool   `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`                                                   // 状态
	Mark            string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                                                     // 备注
}

func (PageInfo) TableName() string { return "ds_page_info" }

type PageVersionInfo struct {
	model.BaseModel
	PageVersionBase
}

type PageVersionBase struct {
	ProductCode string `binding:"required" label:"产品代码" gorm:"column:product_code;size:200;comment:产品代码;" json:"product_code"` // 产品代码
	PortId      string `binding:"required" label:"端口号" gorm:"column:port_id;type:int;comment:端口信息编码;" json:"port_id"`          // 端口信息编码
	Version     string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`                // 版本
	PagePath    string `label:"页面路径" gorm:"column:page_path;size:2000;comment:页面路径;" json:"page_path"`                         // 页面路径
	Mark        string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                                       // 备注
}

func (PageVersionInfo) TableName() string { return "ds_page_version_info" }
