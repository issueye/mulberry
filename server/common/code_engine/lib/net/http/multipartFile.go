package http

import (
	"mime/multipart"

	"github.com/dop251/goja"
)

func NewMultipartFile(runtime *goja.Runtime, file multipart.File) *goja.Object {
	o := runtime.NewObject()

	o.Set("nativeType", file)
	o.Set("read", file.Read)
	o.Set("readAt", file.ReadAt)
	o.Set("seek", file.Seek)
	o.Set("close", file.Close())

	return o
}
