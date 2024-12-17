package model

import (
	"mulberry/internal/host/app/common/model"
	"time"
)

// TaskType 定义任务类型枚举
type TaskType string

const (
	// ScriptTask 脚本任务类型
	ScriptTask TaskType = "script"
	// HTTPTask HTTP任务类型
	HTTPTask TaskType = "http"
)

type ScriptLanguage string

const (
	// JavaScript 脚本语言为JavaScript
	JavaScript ScriptLanguage = "javascript"
	// TypeScript 脚本语言为TypeScript
	TypeScript ScriptLanguage = "typescript"
	// Shell 脚本语言为Shell
	Shell ScriptLanguage = "shell"
)

type Task struct {
	model.BaseModel
	TaskBase
}

type TaskBase struct {
	Name              string     `gorm:"column:name;size:255;not null;comment:名称;" json:"name"`
	Description       string     `gorm:"column:description;size:-1;comment:描述;" json:"description"`
	CronExpression    string     `gorm:"column:cron_expression;size:100;not null;comment:cron 表达式;" json:"cron_expression"`
	TaskType          TaskType   `gorm:"column:task_type;size:50;not null;comment:任务类型;" json:"task_type"`
	Status            int        `gorm:"column:status;not null;default:0;comment:状态;" json:"status"`
	LastExecutionTime *time.Time `gorm:"column:last_execution_time;comment:上一次执行时间;" json:"last_execution_time"`
	NextExecutionTime *time.Time `gorm:"column:next_execution_time;comment:下一次执行时间;" json:"next_execution_time"`
	// 当任务类型是脚本时，以下字段有效
	ScriptLanguage *ScriptLanguage `gorm:"column:script_language;size:50;comment:脚本语言;" json:"script_language"`
	ScriptContent  string          `gorm:"column:script_content;size:-1;comment:脚本内容;" json:"script_content"`
	// 客户端ID， 随机 32 位字符串
	ClientAuthId string `gorm:"column:client_auth_id;size:255;not null;comment:客户端认证ID;" json:"client_auth_id"` // 客户端认证ID
}

// TableName 为Task结构体指定表名
func (Task) TableName() string {
	return "task_info"
}
