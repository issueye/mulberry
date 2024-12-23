package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/google/uuid"
)

func GetNativeType(runtime *goja.Runtime, call *goja.FunctionCall, idx int) interface{} {
	return call.Argument(idx).ToObject(runtime).Get("nativeType").Export()
}

func GetGoType(runtime *goja.Runtime, call *goja.FunctionCall, idx int) goja.Value {
	p := call.Argument(idx).ToObject(runtime)
	protoFunc, ok := goja.AssertFunction(p.Get("getGoType"))
	if !ok {
		panic(runtime.NewTypeError("p%d not have getGoType() function", idx))
	}
	obj, err := protoFunc(p)
	if err != nil {
		panic(runtime.NewGoError(err))
	}
	return obj
}

func InitUtils() {
	require.RegisterNativeModule("std/utils", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("print", func(msg string) {
			fmt.Print(msg)
		})

		o.Set("panic", func(msg string) { panic(msg) })
		o.Set("toString", func(data []byte) string { return string(data) })
		o.Set("uuid", func() string { return uuid.NewString() })

		o.Set("toBase64", func(data string) string {
			return base64.StdEncoding.EncodeToString([]byte(data))
		})

		o.Set("md5", func(data string) string {
			tmp := md5.Sum([]byte(data))
			return hex.EncodeToString(tmp[:])
		})

		o.Set("sha1", func(data string) string {
			tmp := sha1.Sum([]byte(data))
			return hex.EncodeToString(tmp[:])
		})

	})
}
