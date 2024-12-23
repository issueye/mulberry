package code_engine

import (
	"carambola/common/code_engine/lib"
	"fmt"
	"net/http"

	libdb "carambola/common/code_engine/lib/libdb"
	libhttp "carambola/common/code_engine/lib/net/http"

	"github.com/dop251/goja"
	"gorm.io/gorm"
)

func init() {
	lib.InitStrings()  // 字符串
	lib.InitFmt()      // fmt
	lib.InitCrypto()   // 加密
	lib.InitFilepath() // 文件路径
	lib.InitUtils()    //
	lib.InitTime()     // 时间
	lib.InitIni()      // INI
	libhttp.InitHttp() // http
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
