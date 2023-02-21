package main

import (
	"fmt"
	"framework"
	"time"
)

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
	core.Get("foo/:id/ccc/:a", FooControllerHandler)

	bgourp := core.Group("b")
	bgourp.Get("/ccc", func(c *framework.Context) error {
		c.Json("halow world")
		return nil
	})

	dgroup := bgourp.Group("ddd")
	dgroup.Use(func(c *framework.Context) error {
		c.Next()
		fmt.Println("int /ccc/ddd")
		return nil
	})

	dgroup.Get("f", func(c *framework.Context) error {
		time.Sleep(10 * time.Second)
		fmt.Println("int /ccc/ddd/f")
		c.SetStatus(200).Json("/ccc/ddd/fff")
		return nil
	})
}
