package http

import (
	"net/http"

	js "github.com/dop251/goja"
)

func NewCookie(runtime *js.Runtime, cookie *http.Cookie) *js.Object {
	o := runtime.NewObject()

	o.Set("string", cookie.String())
	o.Set("getDomain", cookie.Domain)
	o.Set("getExpires", cookie.Expires.Format("2006-01-02 15:04:05"))
	o.Set("getHttpOnly", cookie.HttpOnly)
	o.Set("getMaxAge", cookie.MaxAge)
	o.Set("getName", cookie.Name)
	o.Set("getPath", cookie.Path)
	o.Set("getRaw", cookie.Raw)
	o.Set("getRawExpires", cookie.RawExpires)
	o.Set("getSecure", cookie.Secure)
	o.Set("getUnparsed", cookie.Unparsed)
	o.Set("getValue", cookie.Value)

	return o
}
