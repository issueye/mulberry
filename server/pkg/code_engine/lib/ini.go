package lib

import (
	"strconv"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	ini "gopkg.in/ini.v1"
)

var iniMap = make(map[string]*ini.File)

func NewIni(runtime *js.Runtime, cfg *ini.File, path string) js.Value {
	o := runtime.NewObject()
	o.Set("getStr", func(sectionStr, key string) string { return cfg.Section(sectionStr).Key(key).String() })
	o.Set("getInt", func(sectionStr, key string) int64 { return cfg.Section(sectionStr).Key(key).MustInt64(-1) })
	o.Set("getBool", func(sectionStr, key string) bool { return cfg.Section(sectionStr).Key(key).MustBool(false) })

	o.Set("getSection", func(sectionStr string) (map[string]string, error) {
		section, err := cfg.GetSection(sectionStr)
		if err != nil {
			return nil, err
		}

		keys := section.KeysHash()
		return keys, nil
	})

	// set
	o.Set("setStr", func(sectionStr, key, value string) { cfg.Section(sectionStr).Key(key).SetValue(value) })
	o.Set("setInt", func(sectionStr, key string, value int64) {
		cfg.Section(sectionStr).Key(key).SetValue(strconv.FormatInt(value, 10))
	})

	o.Set("setBool", func(sectionStr, key string, value bool) {
		cfg.Section(sectionStr).Key(key).SetValue(strconv.FormatBool(value))
	})

	o.Set("save", func(path string) error { return cfg.SaveTo(path) })
	return o
}

func InitIni() {
	require.RegisterNativeModule("go/ini", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("create", func(path string) js.Value {
			iniCfg, err := ini.Load(path)
			if err != nil {
				return nil
			}

			iniMap[path] = iniCfg
			return NewIni(runtime, iniMap[path], path)
		})
	})
}
