package initialize

import (
	"carambola/common/db"
	"carambola/host/app/admin/logic"
	adminModel "carambola/host/app/admin/model"
	taskModel "carambola/host/app/task/model"
	"carambola/host/global"
	"path/filepath"

	"gorm.io/gorm"

	taskLogic "carambola/host/app/task/logic"

	"github.com/nalgeon/redka"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func InitDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "data.db")
	global.DB = db.InitSqlite(path, global.Logger.Sugar())

	InitDATA(global.DB)
}

func InitDATA(db *gorm.DB) {
	db.AutoMigrate(&adminModel.User{})
	db.AutoMigrate(&adminModel.Role{})
	db.AutoMigrate(&adminModel.UserRole{})
	db.AutoMigrate(&adminModel.RoleMenu{})
	db.AutoMigrate(&adminModel.Menu{})

	db.AutoMigrate(&taskModel.Task{})
	db.AutoMigrate(&taskModel.TaskExecutionHistory{})
	db.AutoMigrate(&taskModel.Client{})
	db.AutoMigrate(&taskModel.DatabaseInfo{})

	logic.InitRoles()
	logic.InitRoleMenus()
	logic.InitUserRole()
	logic.InitAdminUser()
	logic.InitMenus()
}

func FreeDB() {
	sqldb, err := global.DB.DB()
	if err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}

	if err = sqldb.Close(); err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}
}

func InitRedkaDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "rdb.db")

	rdb, err := redka.Open(path, nil)
	if err != nil {
		panic(err)
	}

	global.RedkaDB = rdb
}

func InitRegDB() {
	taskLogic.InitTaskDB()
}
