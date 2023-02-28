package env

import (
	"framework"
	"framework/contract"
)

type LsfEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *LsfEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeEnv
}

// Boot will called when the service instantiate
func (provider *LsfEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *LsfEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *LsfEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *LsfEnvProvider) Name() string {
	return contract.EnvKey
}
