package app

import (
	"framework"
	"framework/contract"
)

type LsfAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (app *LsfAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewLsfApp
}

// Boot 启动调用
func (app *LsfAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (app *LsfAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (app *LsfAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, app.BaseFolder}
}

// Name 获取字符串凭证
func (app *LsfAppProvider) Name() string {
	return contract.AppKey
}
