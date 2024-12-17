package lib

import (
	"mulberry/pkg/utils"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitCrypto() {
	require.RegisterNativeModule("go/crypto", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("md5", func(value any) any {
			switch data := value.(type) {
			case []byte:
				return utils.MD5V(utils.Bytes2String(data))
			case string:
				return utils.MD5V(data)
			default:
				return value
			}
		})

		o.Set("sha1", func(value any) any {
			switch data := value.(type) {
			case []byte:
				return utils.Sha1(utils.Bytes2String(data))
			case string:
				return utils.Sha1(data)
			default:
				return value
			}
		})
	})
}
