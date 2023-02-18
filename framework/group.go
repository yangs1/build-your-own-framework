package framework

import "strings"

// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
	AddRouter(string, string, ControllerHandler)
	Use(...ControllerHandler)
	// Group 实现嵌套group
	Group(string) IGroup
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core  // 指向core结构
	parent *Group //指向上一个Group，如果有的话
	prefix string // 这个group的通用前缀

	middlewares []ControllerHandler // 中间件
}

// NewGroup 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// Get 实现Get方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	g.AddRouter("GET", uri, handler)
}

// Post 实现Post方法
func (g *Group) Post(uri string, handler ControllerHandler) {
	g.AddRouter("POST", uri, handler)
}

// Put 实现Put方法
func (g *Group) Put(uri string, handler ControllerHandler) {
	g.AddRouter("PUT", uri, handler)
}

// Delete 实现Delete方法
func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.AddRouter("DELETE", uri, handler)
}

// AddRouter 实现AddRouter方法
func (g *Group) AddRouter(method string, uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + "/" + strings.Trim(uri, "/")
	g.core.AddRouter(method, uri, handler, g.middlewares...)
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + "/" + g.prefix
}

// 实现 Group 方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, strings.Trim(uri, "/"))
	cgroup.parent = g
	return cgroup
}

//注册中间件
func (c *Group) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}
