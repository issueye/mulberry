package model

import (
	"mulberry/internal/host/app/common/model"
)

// PortInfo
// 端口信息
type GzipFilterInfo struct {
	model.BaseModel
	GzipFilterBase
}

type GzipFilterBase struct {
	PortId       string `gorm:"column:port_id;type:int;comment:端口信息编码;" json:"portId"`            // 端口信息编码
	MatchContent string `gorm:"column:match_content;size:2000;comment:匹配内容;" json:"matchContent"` // 匹配内容
	MatchType    uint   `gorm:"column:match_type;type:int;comment:匹配模式;" json:"matchType"`        // 匹配模式
	Mark         string `gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                    // 备注
}

func (mod *GzipFilterInfo) Copy(data *GzipFilterBase) {
	mod.PortId = data.PortId
	mod.MatchContent = data.MatchContent
	mod.MatchType = data.MatchType
	mod.Mark = data.Mark
}

// TableName
// 表名称
func (GzipFilterInfo) TableName() string {
	return "bridge_gzip_filter_info"
}
