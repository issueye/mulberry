package lib

import (
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitFile() {
	require.RegisterNativeModule("go/file", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("write", func(filename string, data string) error {
			err := os.WriteFile(filename, []byte(data), 0666)
			if err != nil {
				return err
			}
			return nil
		})

		o.Set("read", func(filename string) (string, error) {
			data, err := os.ReadFile(filename)
			if err != nil {
				return "", err
			}
			return string(data), nil
		})
	})
}
