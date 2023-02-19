package main

import (
	"fmt"
	"framework"
	"time"
)

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)

	bgourp := core.Group("b")
	bgourp.Get("/ccc", func(c *framework.Context) error {
		return c.Json(200, "halow world")
	})

	dgroup := bgourp.Group("ddd")
	dgroup.Use(func(c *framework.Context) error {
		c.Next()
		fmt.Println("int /ccc/ddd")
		return nil
	})

	dgroup.Get("f", func(c *framework.Context) error {
		fmt.Println("int /ccc/ddd/f")
		time.Sleep(10 * time.Second)
		return c.Json(200, "/ccc/ddd/fff")
	})
}
