package lib

import (
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitFilepath() {
	require.RegisterNativeModule("go/filepath", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("abs", filepath.Abs)
		o.Set("join", filepath.Join)
		o.Set("ext", filepath.Ext)
		o.Set("rel", filepath.Rel)
		o.Set("clean", filepath.Clean)
		o.Set("split", filepath.Split)
		o.Set("dir", filepath.Dir)
		o.Set("base", filepath.Base)
	})
}
