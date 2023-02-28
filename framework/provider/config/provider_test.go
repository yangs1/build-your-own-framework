package config

import (
	"framework"
	"framework/contract"
	"framework/provider/app"
	"framework/provider/env"
	"log"
	"testing"
)

func TestHadeConfig_Normal(t *testing.T) {

	c := framework.NewLsfContainer()
	c.Bind(&app.LsfAppProvider{BaseFolder: "/Users/Yang/Workspace/Code/Goland/golsf"})
	c.Bind(&env.LsfEnvProvider{})
	c.Bind(&ConfigProvider{})

	conf := c.MustMake(contract.ConfigKey).(contract.Config)

	log.Println(conf.GetString("database.mysql.hostname"))

	maps2 := conf.GetStringMapString("databse.mysql")
	log.Println(maps2["timeout"]) //""

	type Mysql struct {
		Hostname string
		Username string
	}
	ms := &Mysql{}
	conf.Load("database.mysql", ms)
	log.Println(ms)

}
