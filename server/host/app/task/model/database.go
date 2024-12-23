package model

import (
	"carambola/host/common/model"
	"encoding/json"
)

// DatabaseType 定义数据库类型枚举
type DatabaseType string

const (
	// MySQL 数据库类型为MySQL
	MySQL DatabaseType = "mysql"
	// PostgreSQL 数据库类型为PostgreSQL
	PostgreSQL DatabaseType = "postgresql"
	// SQLite 数据库类型为SQLite
	SQLite DatabaseType = "sqlite"
	// Oracle 数据库类型为Oracle
	Oracle DatabaseType = "oracle"
	// SQLServer 数据库类型为SQLServer
	SQLServer DatabaseType = "sqlserver"
)

type DatabaseInfo struct {
	model.BaseModel
	DatabaseBase
}

type DatabaseBase struct {
	Name     string       `gorm:"column:name;size:255;not null;comment:名称;" json:"name"`
	DBType   DatabaseType `gorm:"column:db_type;size:255;not null;comment:数据库类型;" json:"db_type"`
	Host     string       `gorm:"column:host;size:255;not null;comment:主机;" json:"host"`
	Port     int          `gorm:"column:port;not null;comment:端口;" json:"port"`
	Username string       `gorm:"column:username;size:255;not null;comment:用户名;" json:"username"`
	Password string       `gorm:"column:password;size:255;not null;comment:密码;" json:"password"`
	Database string       `gorm:"column:database;size:255;not null;comment:数据库;" json:"database"`
	Schema   string       `gorm:"column:schema;size:255;not null;comment:模式;" json:"schema"`
	Path     string       `gorm:"column:path;size:255;not null;comment:路径;" json:"path"`
}

// TableName 为Task结构体指定表名
func (DatabaseInfo) TableName() string {
	return "database_info"
}

func (tb *DatabaseInfo) ToJson() string {
	data, err := json.Marshal(tb)
	if err != nil {
		return ""
	}

	return string(data)
}

func (DatabaseInfo) FromJson(value string) *DatabaseInfo {

	db := &DatabaseInfo{}

	err := json.Unmarshal([]byte(value), db)
	if err != nil {
		return nil
	}

	return db
}
