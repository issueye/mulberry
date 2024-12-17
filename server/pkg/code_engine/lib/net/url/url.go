package url

import (
	"encoding/json"
	"net/url"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func NewURL(runtime *goja.Runtime, u *url.URL) *goja.Object {
	o := runtime.NewObject()

	o.Set("getForceQuery", u.ForceQuery)
	o.Set("getFragment", u.Fragment)
	o.Set("getHost", u.Host)
	o.Set("getOpaque", u.Opaque)
	o.Set("getPath", u.Path)
	o.Set("getRawPath", u.RawPath)
	o.Set("getRawQuery", u.RawQuery)
	o.Set("getScheme", u.Scheme)
	o.Set("getPort", u.Port())

	return o
}

func InitUrl() {
	require.RegisterNativeModule("url", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("parse", func(rawUrl string) goja.Value {
			u, err := url.Parse(rawUrl)
			if err != nil {
				return nil
			}

			return NewURL(runtime, u)
		})

		o.Set("queryEscape", url.QueryEscape)

		o.Set("queryUnescape", func(arg string) string {
			str, err := url.QueryUnescape(arg)
			if err != nil {
				return ""
			}
			return str
		})

		o.Set("parseRequestURI", func(rawUrl string) goja.Value {
			mUrl, err := url.ParseRequestURI(rawUrl)
			if err != nil {
				return nil
			}
			return NewURL(runtime, mUrl)
		})

		o.Set("parseQuery", func(query string) string {
			values, err := url.ParseQuery(query)
			if err != nil {
				return "{}"
			}
			if len(values) > 0 {
				val, err := json.Marshal(values)
				if err != nil {
					return "{}"
				}
				return string(val)
			}

			return "{}"
		})

		o.Set("newValues", func() goja.Value {
			return NewValues(runtime, make(url.Values))
		})
	})
}
