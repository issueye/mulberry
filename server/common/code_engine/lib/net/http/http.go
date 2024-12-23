package http

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitHttp() {
	require.RegisterNativeModule("http", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("request", func(method string, rawUrl string, headers map[string]any, body string, timeout int) map[string]any {
			req, err := http.NewRequest(method, rawUrl, strings.NewReader(body))
			if err != nil {
				return nil
			}

			for k, v := range headers {
				if str, ok := v.(string); ok {
					req.Header.Set(k, str)
				} else {
					return nil
				}
			}

			client := &http.Client{}
			client.Timeout = time.Duration(timeout) * time.Millisecond
			res, err := client.Do(req)
			if err != nil {
				return nil
			}
			defer res.Body.Close()
			datas, err := io.ReadAll(res.Body)
			if err != nil {
				return nil
			}

			headerObj := runtime.NewObject()
			for k, v := range res.Header {
				headerObj.Set(k, v)
			}

			return map[string]any{
				"body":   datas,
				"header": headerObj,
			}
		})
	})
}
