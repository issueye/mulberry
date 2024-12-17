package db

import (
	"github.com/dop251/goja"
	"gorm.io/gorm"
)

func NewTx(runtime *goja.Runtime, tx *gorm.DB) goja.Value {
	o := runtime.NewObject()

	o.Set("commit", func() { tx.Commit() })
	o.Set("rollback", func() { tx.Rollback() })

	o.Set("exec", func(sqlStr string) (int64, error) {
		result := tx.Exec(sqlStr)
		if result.Error != nil {
			return -1, result.Error
		}

		// 获取返回
		return result.RowsAffected, nil
	})

	o.Set("query", func(sqlStr string) ([]map[string]any, error) {
		// 查询数据
		result := tx.Raw(sqlStr)

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

	return o
}
