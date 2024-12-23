package model

import (
	"carambola/host/common/model"
)

// TaskExecutionHistory 任务执行历史结构体
type TaskExecutionHistory struct {
	model.BaseModel
	TaskExecutionHistoryBase
}

type TaskExecutionHistoryBase struct {
	TaskID            string `gorm:"column:task_id;not null" json:"task_id"`                // 任务ID
	ClientID          string `gorm:"column:client_id;not null" json:"client_id"`            // 客户端ID
	ExecutionID       string `gorm:"column:execution_id;not null" json:"execution_id"`      // 执行ID
	StartTime         string `gorm:"column:start_time" json:"start_time"`                   // 开始时间
	EndTime           string `gorm:"column:end_time" json:"end_time"`                       // 结束时间
	NextExecutionTime string `gorm:"column:next_execution_time" json:"next_execution_time"` // 下次执行时间
	Status            int    `gorm:"column:status" json:"status"`                           // 状态
	Result            string `gorm:"column:result" json:"result"`                           // 结果
}

// TableName 为TaskExecutionHistory结构体指定表名
func (TaskExecutionHistory) TableName() string { return "task_execution_history" }
