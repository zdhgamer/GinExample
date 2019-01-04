package main

import (
	"net/http"
	"fmt"
	"../setting"
	"../routers"
)

func main() {
	// 注册一个默认的路由器
	router := routers.InitRouter()
	// 绑定端口是8000
	fmt.Println(setting.HttpPort)
	s := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.HttpPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	s.ListenAndServe()
}
