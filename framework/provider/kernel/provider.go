package kernel

import (
	"framework"
	"framework/contract"
)

// KernelProvider 提供web引擎
type KernelProvider struct {
	Core *framework.Core
}

// Register 注册服务提供者
func (provider *KernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewKernelService
}

func (provider *KernelProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer 引擎的初始化我们希望开始就进行初始化
func (provider *KernelProvider) IsDefer() bool {
	return false
}

// Params 参数就是一个 HttpEngine
func (provider *KernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Core}
}

// Name 提供凭证
func (provider *KernelProvider) Name() string {
	return contract.KernelKey
}
