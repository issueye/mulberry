package url

import (
	"net/url"

	"github.com/dop251/goja"
)

func NewValues(runtime *goja.Runtime, values url.Values) *goja.Object {
	o := runtime.NewObject()

	o.Set("del", values.Del)
	o.Set("add", values.Add)
	o.Set("encode", values.Encode)
	o.Set("get", values.Get)
	o.Set("gets", func(key string) []string { return values[key] })
	o.Set("getAll", map[string][]string(values))
	o.Set("set", values.Set)
	return o
}
