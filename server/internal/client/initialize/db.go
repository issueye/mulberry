package initialize

import (
	"context"
	"mulberry/internal/global"
	"mulberry/internal/host/app/task/model"
	"mulberry/pkg/code_engine"
	"mulberry/pkg/db"
	"mulberry/pkg/ipc/grpc/pb"
)

func InitDB() {
	list, err := global.IpcClient.RDBHelper().ListRange(context.Background(), &pb.RDBStartEnd{Key: global.DB_Key, Start: 0, End: -1})
	if err != nil {
		global.Logger.Sugar().Errorf("获取数据信息失败: %s", err.Error())
		return
	}

	global.Logger.Sugar().Infof("获取数据信息成功, 数量: %d", len(list.Values))
	InitData(list.Values)
}

// 初始化数据
func InitData(list []string) {
	for _, data := range list {
		info := model.DatabaseInfo{}.FromJson(data)
		switch info.DBType {
		case model.SQLite:
			InitSqlite(info)
		case model.MySQL:
			InitMysql(info)
		case model.SQLServer:
			InitSqlServer(info)
		case model.Oracle:
			InitOracle(info)
		}
	}
}

func InitSqlite(info *model.DatabaseInfo) {
	gdb := db.InitSqlite(info.Path, global.Logger.Sugar())
	global.DBMap[info.Name] = gdb
	code_engine.InitDB(info.Name, gdb)
}

func InitMysql(info *model.DatabaseInfo) {
	gdb := db.InitMysql(
		&db.Config{
			Username: info.Username,
			Password: info.Password,
			Host:     info.Host,
			Database: info.Database,
			Port:     info.Port,
		},
		global.Logger.Sugar(),
	)
	global.DBMap[info.Name] = gdb
	code_engine.InitDB(info.Name, gdb)
}

func InitSqlServer(info *model.DatabaseInfo) {
	gdb := db.InitSqlServer(
		&db.Config{
			Username: info.Username,
			Password: info.Password,
			Host:     info.Host,
			Database: info.Database,
			Port:     info.Port,
		},
		global.Logger.Sugar(),
	)
	global.DBMap[info.Name] = gdb
	code_engine.InitDB(info.Name, gdb)
}

func InitOracle(info *model.DatabaseInfo) {
	gdb := db.InitOracle(
		&db.Config{
			Username: info.Username,
			Password: info.Password,
			Host:     info.Host,
			Database: info.Database,
			Port:     info.Port,
		},
		global.Logger.Sugar(),
	)
	global.DBMap[info.Name] = gdb
	code_engine.InitDB(info.Name, gdb)
}
