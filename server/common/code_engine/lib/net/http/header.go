package http

import (
	"bytes"
	"net/http"

	"github.com/dop251/goja"
)

func NewHeader(runtime *goja.Runtime, header http.Header) *goja.Object {
	o := runtime.NewObject()

	o.Set("add", header.Add)
	o.Set("del", header.Del)
	o.Set("get", header.Get)
	o.Set("gets", func(key string) []string { return header[key] })
	o.Set("set", header.Set)
	o.Set("getRaw", func() string {
		byteBuf := &bytes.Buffer{}
		header.Write(byteBuf)
		return string(byteBuf.Bytes())
	})

	return o
}
