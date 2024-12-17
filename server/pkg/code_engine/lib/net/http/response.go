package http

import (
	"net/http"

	"github.com/dop251/goja"
)

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	o := runtime.NewObject()
	o.Set("header", func() goja.Value {
		return NewHeader(runtime, w.Header())
	})

	o.Set("write", func(data string) (int, error) {
		return w.Write([]byte(data))
	})

	o.Set("writeHeader", func(n int) { w.WriteHeader(n) })

	o.Set("setCookie", func(name string, value string, path string, maxAge int, httpOnly bool) {
		cookie := &http.Cookie{}
		cookie.Name = name
		cookie.Value = value
		cookie.Path = path
		cookie.MaxAge = maxAge
		cookie.HttpOnly = httpOnly
		http.SetCookie(w, cookie)
	})

	o.Set("nativeType", w)

	return o
}
