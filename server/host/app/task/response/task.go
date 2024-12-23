package response

type TaskLogResponse struct {
	TaskID            string `gorm:"column:task_id" json:"task_id"`                         // 任务ID
	TaskName          string `gorm:"column:task_name" json:"task_name"`                     // 任务名称
	ClientID          string `gorm:"column:client_id" json:"client_id"`                     // 客户端ID
	ClientName        string `gorm:"column:client_name" json:"client_name"`                 // 客户端名称
	ExecutionID       string `gorm:"column:execution_id" json:"execution_id"`               // 执行ID
	StartTime         string `gorm:"column:start_time" json:"start_time"`                   // 开始时间
	EndTime           string `gorm:"column:end_time" json:"end_time"`                       // 结束时间
	NextExecutionTime string `gorm:"column:next_execution_time" json:"next_execution_time"` // 下次执行时间
	Status            int    `gorm:"column:status" json:"status"`                           // 状态
	Result            string `gorm:"column:result" json:"result"`                           // 结果
}
