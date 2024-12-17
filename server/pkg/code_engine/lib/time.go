package lib

import (
	"time"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitTime() {
	require.RegisterNativeModule("go/time", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		// 睡眠
		o.Set("sleep", func(d int64) {
			<-time.After(time.Duration(d) * time.Millisecond)
		})

		// 当前时间字符串
		o.Set("nowString", time.Now().Format("2006-01-02 15:04:05"))
		o.Set("nowDate", time.Now().Format("2006-01-02"))
		o.Set("nowYear", time.Now().Format("2006"))
		o.Set("nowMonth", time.Now().Format("01"))
		o.Set("nowDay", time.Now().Format("02"))
		o.Set("nowHour", time.Now().Format("15"))
		o.Set("nowMinute", time.Now().Format("04"))
		o.Set("nowSecond", time.Now().Format("05"))
	})
}
