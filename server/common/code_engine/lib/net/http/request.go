package http

import (
	"carambola/common/code_engine/lib/net/url"
	"io"
	"net/http"

	"github.com/dop251/goja"
)

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	o := runtime.NewObject()

	o.Set("getContentLength", r.ContentLength)
	o.Set("getMethod", r.Method)
	o.Set("getHost", r.Host)
	o.Set("getBody", r.Body)
	o.Set("getHeader", NewHeader(runtime, r.Header))
	o.Set("getHeaders", map[string][]string(r.Header))
	o.Set("getUri", r.RequestURI)
	o.Set("getUrl", url.NewURL(runtime, r.URL))
	o.Set("getRemoteAddr", r.RemoteAddr)
	o.Set("getForm", url.NewValues(runtime, r.Form))
	o.Set("formValue", r.FormValue)
	o.Set("userAgent", r.UserAgent())
	o.Set("parseForm", r.ParseForm)

	o.Set("formFile", func(key string) map[string]any {
		file, fileHeader, err := r.FormFile(key)
		if err != nil {
			return nil
		}

		return map[string]any{
			"file":   NewMultipartFile(runtime, file),
			"name":   fileHeader.Filename,
			"header": map[string][]string(fileHeader.Header),
		}
	})

	o.Set("parseMultipartForm", r.ParseMultipartForm)
	o.Set("cookie", func(name string) goja.Value {
		c, err := r.Cookie(name)
		if err != nil {
			return nil
		}
		return NewCookie(runtime, c)
	})

	o.Set("cookies", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.Cookies())
	})

	o.Set("getRawBody", func(call goja.FunctionCall) string {
		bts, err := io.ReadAll(r.Body)
		if err != nil {
			return ""
		}

		return string(bts)
	})

	o.Set("getBodyString", func(call goja.FunctionCall) string {
		bts, err := io.ReadAll(r.Body)
		if err != nil {
			return "{}"
		}

		return string(bts)
	})

	return o
}
