package framework

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: make(map[string]*Tree),
	}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println("core.serveHTTP")
	ctx := NewContext(w, r)

	routeNode := c.FindRouteByRequest(r)

	if routeNode == nil {
		// 如果没有找到，这里打印日志
		ctx.SetStatus(404).Json("not found")
		return
	}

	handler := routeNode.handler
	middlewares := routeNode.middlewares
	params := routeNode.parseParamsFromEndNode(r.URL.Path)
	// 注入参数
	ctx.SetParams(params)
	// 实现 pipeline 链
	ctx.setHandlers(append(middlewares, handler))

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		fmt.Println("number of goroutines:", runtime.NumGoroutine())
		return
	}
	fmt.Println("number of goroutines:", runtime.NumGoroutine())

}

func (c *Core) FindRouteByRequest(request *http.Request) *node {
	url := request.URL.Path
	method := request.Method

	upperMethod := strings.ToUpper(method)

	//log.Println(upperMethod)
	//log.Println(c.router)
	if methodTree, ok := c.router[upperMethod]; ok {
		return methodTree.FindHandler(url)
	}

	return nil
}

func (c *Core) AddRouter(method string, url string, handler ControllerHandler, middlewares ...ControllerHandler) {
	upperMethod := strings.ToUpper(method)
	_, ok := c.router[upperMethod]
	if !ok {
		c.router[upperMethod] = NewTree()
	}
	err := c.router[upperMethod].AddRouter(url, handler, append(c.middlewares, middlewares...))
	if err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.AddRouter("GET", url, handler)
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.AddRouter("POST", url, handler)
}

func (c *Core) Put(url string, handler ControllerHandler) {
	c.AddRouter("PUT", url, handler)
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	c.AddRouter("DELETE", url, handler)
}

//注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
