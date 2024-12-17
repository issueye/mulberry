package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/google/uuid"
)

func GetNativeType(runtime *js.Runtime, call *js.FunctionCall, idx int) interface{} {
	return call.Argument(idx).ToObject(runtime).Get("nativeType").Export()
}

func GetGoType(runtime *js.Runtime, call *js.FunctionCall, idx int) js.Value {
	p := call.Argument(idx).ToObject(runtime)
	protoFunc, ok := js.AssertFunction(p.Get("getGoType"))
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
	require.RegisterNativeModule("std/utils", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
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
