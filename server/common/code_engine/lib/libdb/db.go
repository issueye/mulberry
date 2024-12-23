package libdb

import (
	"database/sql"
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"gorm.io/gorm"
)

type sqlResult struct {
	LastInsertId int64
	RowsAffected int64
}

func NewExecResult(runtime *goja.Runtime, r sqlResult) goja.Value {
	o := runtime.NewObject()
	o.Set("rowsAffected", r.RowsAffected)
	return o
}

// MakeData
// 生成数据
func MakeData(rows *sql.Rows) ([]map[string]any, error) {
	data := make([]map[string]interface{}, 0)
	cols, err := rows.Columns()
	if err != nil {
		fmt.Printf("查询失败，失败原因：%s\n", err.Error())
		return nil, err
	}

	for rows.Next() {
		columns := make([]any, len(cols))
		pointers := make([]any, len(cols))

		for i := range columns {
			pointers[i] = &columns[i]
		}
		err := rows.Scan(pointers...)
		if err != nil {
			fmt.Printf("绑定数据失败，失败原因：%s\n", err.Error())
			return nil, err
		}

		oneData := make(map[string]any)
		for i, v := range cols {
			valueP := pointers[i].(*any)
			value := *valueP
			oneData[v] = value
		}
		data = append(data, oneData)
	}
	return data, nil
}

type RegNativeFunc = func(name string, loader require.ModuleLoader)

func InitDB(moduleName string, gdb *gorm.DB) {
	require.RegisterNativeModule(moduleName, func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		// query 查询
		o.Set("query", func(sqlStr string) ([]map[string]any, error) {
			// 查询数据
			result := gdb.Raw(sqlStr)

			if result.Error != nil {
				return nil, result.Error
			}

			rows, err := result.Rows()
			if err != nil {
				return nil, err
			}

			// 生成数据
			return MakeData(rows)
		})

		// exec 执行语句 增删改
		o.Set("exec", func(sqlStr string) (int64, error) {
			result := gdb.Exec(sqlStr)
			if result.Error != nil {
				return -1, result.Error
			}

			// 获取返回
			return result.RowsAffected, nil
		})

		// 事务
		// begin 开启事务
		o.Set("begin", func() goja.Value {
			tx := gdb.Begin()
			if tx.Error != nil {
				return nil
			}

			return NewTx(runtime, tx)
		})
	})
}
