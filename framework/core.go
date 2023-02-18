package framework

import (
	"log"
	"net/http"
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
	log.Println("core.serveHTTP")
	ctx := NewContext(w, r)

	router := c.FindRouteByRequest(r)

	if router == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	url := request.URL.Path
	method := request.Method

	upperMethod := strings.ToUpper(method)

	log.Println(upperMethod)
	log.Println(c.router)
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
