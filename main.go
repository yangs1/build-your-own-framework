package main

import (
	"context"
	"fmt"
	"framework"
	"framework/middlewares"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()

	core.Use(middlewares.Recovery())
	core.Use(middlewares.Cost())
	//core.Use(middlewares.Timeout(10 * time.Second))

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	// 这个goroutine是启动服务的goroutine
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	fmt.Println("666777")
	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	log.Println(timeoutCtx)

	defer cancel()

	err := server.Shutdown(timeoutCtx)
	log.Println(err)
	if err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("666")
}
