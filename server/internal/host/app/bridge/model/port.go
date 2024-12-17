package model

import (
	"mulberry/internal/host/app/common/model"
)

// PortInfo
// 端口信息
type PortInfo struct {
	model.BaseModel
	PortBase
}

type PortBase struct {
	Port            int    `gorm:"column:port;type:int;comment:端口号;" json:"port"`                              // 端口号
	State           bool   `gorm:"column:state;type:int;comment:状态 0 停用 1 启用;" json:"state"`                   // 状态
	IsTLS           bool   `gorm:"column:is_tls;type:int;comment:是否https;" json:"isTLS"`                       // 是否证书加密
	CerId           string `gorm:"column:cert_id;size:100;comment:编码;" json:"certId"`                          // 证书编码
	UseGzip         bool   `gorm:"column:use_gzip;type:int;comment:使用GZIP 0 停用 1 启用;" json:"useGzip"`          // 使用GZIP 0 停用 1 启用
	PageCount       int    `gorm:"column:page_count;type:int;comment:页面总数;" json:"pageCount"`                  // 页面总数
	RuleCount       int    `gorm:"column:rule_count;type:int;comment:规则总数;" json:"ruleCount"`                  // 规则总数
	GzipFilterCount int    `gorm:"column:gzip_filter_count;type:int;comment:GZIP过滤总数;" json:"gzipFilterCount"` // GZIP过滤总数
	Mark            string `gorm:"column:mark;size:-1;comment:备注;" json:"mark"`                                // 备注
}

func (mod *PortInfo) Copy(data *PortBase) {
	mod.Port = data.Port
	mod.IsTLS = data.IsTLS
	mod.CerId = data.CerId
	mod.State = data.State
	mod.UseGzip = data.UseGzip
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (PortInfo) TableName() string { return "bridge_port_info" }
