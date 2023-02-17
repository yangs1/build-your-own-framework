package main

import "framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
