package lib

import (
	"fmt"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitFmt() {
	require.RegisterNativeModule("go/fmt", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("sprintf", func(format string, args ...any) string {
			return fmt.Sprintf(format, args...)
		})

		o.Set("printf", fmt.Printf)
		o.Set("println", fmt.Println)
		o.Set("print", fmt.Print)
	})
}
