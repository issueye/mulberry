package code_engine

import (
	"fmt"
	"mulberry/pkg/code_engine/lib"
	"net/http"

	libdb "mulberry/pkg/code_engine/lib/db"
	libhttp "mulberry/pkg/code_engine/lib/net/http"

	"github.com/dop251/goja"
	"gorm.io/gorm"
)

func init() {
	lib.InitStrings()  // 字符串
	lib.InitFmt()      // fmt
	lib.InitCrypto()   // 加密
	lib.InitFilepath() // 文件路径
	lib.InitUtils()
	lib.InitTime()

}

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	return libhttp.NewRequest(runtime, r)
}

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	return libhttp.NewResponse(runtime, w)
}

func InitDB(name string, gdb *gorm.DB) {
	fmt.Printf("注册数据库[%s]到脚本引擎中\n", name)
	libdb.InitDB(name, gdb)
}
