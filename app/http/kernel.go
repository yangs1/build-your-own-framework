package http

import (
	"framework"
	"framework/middlewares"
)

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine() (*framework.Core, error) {
	// 默认启动一个Web引擎
	r := framework.New()
	// 初始化中间件
	r.Use(middlewares.Recovery(), middlewares.Cost())

	// 业务绑定路由操作
	registerRouter(r)
	// 返回绑定路由后的Web引擎
	return r, nil
}
