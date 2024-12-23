package model

import "carambola/host/common/model"

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

func (GzipFilterInfo) TableName() string { return "ds_gzip_filter_info" }
