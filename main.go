package main

import (
	"framework"
	"framework/provider/app"
	"framework/provider/config"
	"framework/provider/kernel"
	"golsf/app/console"
	"golsf/app/http"
)

func main() {

	// 初始化服务容器
	container := framework.NewLsfContainer()
	// 绑定App服务提供者
	container.Bind(&app.LsfAppProvider{})
	container.Bind(&config.ConfigProvider{})

	if r, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.KernelProvider{Core: r})
	}

	// 运行root命令
	console.RunCommand(container)
	/*core := framework.NewCore()

	// 绑定具体的服务
	core.Bind(&app.LsfAppProvider{})

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
	log.Println("ok")*/
}
