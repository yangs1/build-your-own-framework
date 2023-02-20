package main

import (
	"framework"
)

func FooControllerHandler(c *framework.Context) error {
	json := map[string]string{"sd": "123"}

	c.SetStatus(200).Json(json)
	return nil
}
