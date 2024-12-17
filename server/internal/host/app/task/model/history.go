package model

import (
	"mulberry/internal/host/app/common/model"

	"gorm.io/gorm"
)

// TaskExecutionHistory 任务执行历史结构体
type TaskExecutionHistory struct {
	model.BaseModel
	TaskID          uint            `gorm:"column:task_id;not null" json:"task_id"`
	StartTime       gorm.DeletedAt  `gorm:"column:start_time" json:"start_time"`
	EndTime         *gorm.DeletedAt `gorm:"column:end_time" json:"end_time"`
	ExecutionStatus int             `gorm:"not null" json:"execution_status"`
	ErrorMessage    string          `gorm:"column:error_message;size:-1" json:"error_message"`
	ExecutionLog    string          `gorm:"column:execution_log;size:-1" json:"execution_log"`
}

// TableName 为TaskExecutionHistory结构体指定表名
func (TaskExecutionHistory) TableName() string {
	return "task_execution_history"
}
