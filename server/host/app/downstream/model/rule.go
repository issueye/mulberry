package model

import "carambola/host/common/model"

type RuleInfo struct {
	model.BaseModel
	RuleBase
}

type RuleBase struct {
	Name        string `gorm:"column:name;size:300;comment:匹配路由名称;" json:"name"`                                              // 匹配路由名称
	TargetId    string `gorm:"column:target_id;size:300;comment:目标路由;" json:"target_id"`                                      // 目标地址编码
	TargetRoute string `gorm:"column:target_route;size:300;comment:目标路由;" json:"target_route"`                                // 目标路由
	PortId      string `gorm:"column:port_id;size:100;comment:端口信息编码;" json:"port_id"`                                        // 端口信息编码
	Method      string `gorm:"column:method;size:100;comment:请求方法;" json:"method"`                                            // 请求方法
	MatchType   uint   `gorm:"column:match_type;type:int;comment:匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配;" json:"match_type"` // 匹配模式 0 所有内容匹配 1 正则匹配 2 包含匹配 3 header 匹配
	Mark        string `gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                                                 // 备注
}

func (RuleInfo) TableName() string { return "dz_rule_info" }
