package main

import (
	"carambola/client/initialize"
	_ "carambola/client/winappres"
	"flag"
)

var (
	ClientID = flag.String("n", "", "client id")
)

func main() {
	// 接收一个 -n 的参数，用于设置当前客户端的ID
	flag.Parse()
	if *ClientID == "" {
		panic("client id is required")
	}

	initialize.Init(*ClientID)

	// 阻塞主线程
	select {}
}
