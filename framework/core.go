package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: make(map[string]ControllerHandler),
	}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(w, r)

	router := c.router["foo"]
	if router == nil {
		return
	}

	router(ctx)
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}
