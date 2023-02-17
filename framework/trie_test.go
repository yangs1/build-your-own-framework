package framework

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	nt := NewTree()
	nt.AddRouter("/foo", func(c *Context) error {
		fmt.Println("666")
		return nil
	})

	nt.FindHandler("/foo")
}
