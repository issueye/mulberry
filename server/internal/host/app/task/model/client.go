package model

import (
	"mulberry/internal/host/app/common/model"
	"time"
)

type Client struct {
	model.BaseModel
	ClientBase
}

type ClientBase struct {
	Name         string     `gorm:"column:name;size:255;not null;comment:名称;" json:"name"`                          // 名称
	ClientAuthId string     `gorm:"column:client_auth_id;size:255;not null;comment:客户端认证ID;" json:"client_auth_id"` // 客户端认证ID
	Status       int        `gorm:"column:status;not null;default:0;comment:状态;" json:"status"`                     // 状态
	StartAt      *time.Time `gorm:"column:start_at;comment:开启客户端时间;" json:"start_at"`                               // 开启客户端时间
	Pid          int        `gorm:"column:pid;not null;comment:进程ID;" json:"pid"`                                   // 进程ID
	Remark       string     `gorm:"column:remark;size:255;not null;comment:备注;" json:"remark"`                      // 备注
}

func (Client) TableName() string {
	return "task_client_info"
}
