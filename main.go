package main

import (
	"context"
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
		log.Println(server.ListenAndServe())
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("ok")
}
