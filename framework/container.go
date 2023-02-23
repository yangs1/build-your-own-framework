package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool
	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type LsfContainer struct {
	Container
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// NewHadeContainer 创建一个服务容器
func NewLsfContainer() *LsfContainer {
	return &LsfContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (container *LsfContainer) PrintfProvider() []string {
	ret := []string{}
	for _, provider := range container.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (container *LsfContainer) Bind(provider ServiceProvider) error {
	container.lock.Lock()
	defer container.lock.Unlock()

	key := provider.Name()
	container.providers[key] = provider

	if provider.IsDefer() {
		// 实例化方法
		instance, err := container.newInstance(provider, nil)

		if err != nil {
			return errors.New(err.Error())
		}
		container.instances[key] = instance
	}

	return nil
}

func (container *LsfContainer) IsBind(key string) bool {
	return container.findServiceProvider(key) != nil
}

func (container *LsfContainer) findServiceProvider(key string) ServiceProvider {
	container.lock.RLock()
	defer container.lock.RUnlock()
	if sp, ok := container.providers[key]; ok {
		return sp
	}
	return nil
}

func (container *LsfContainer) Make(key string) (interface{}, error) {
	return container.make(key, nil, false)
}

func (container *LsfContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return container.make(key, params, true)
}

func (container *LsfContainer) MustMake(key string) interface{} {
	serv, err := container.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (container *LsfContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	container.lock.RLock()
	defer container.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := container.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return container.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := container.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := container.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	container.instances[key] = inst
	return inst, nil
}

func (container *LsfContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(container); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(container)
	}
	method := sp.Register(container)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}
