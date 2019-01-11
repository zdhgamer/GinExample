package main

import (
	"net/http"
	"fmt"
	"../setting"
	"../routers"
	"os"
	"os/signal"
	"log"
	"context"
	"time"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler:router,
		ReadTimeout:setting.ServerSetting.ReadTimeout,
		WriteTimeout:setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	err := s.ListenAndServe()
	if err!=nil {
		fmt.Printf("start server err:%v",err)
		panic(s)
	}

	//优雅的关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exist")
}
