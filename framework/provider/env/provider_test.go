package env

import (
	"framework"
	"framework/contract"
	"framework/provider/app"
	"log"
	"testing"
)

func TestHadeEnvProvider(t *testing.T) {

	c := framework.NewLsfContainer()
	sp := &app.LsfAppProvider{}

	err := c.Bind(sp)
	log.Println(err)

	sp2 := &LsfEnvProvider{}
	err = c.Bind(sp2)
	log.Println(err)

	envServ := c.MustMake(contract.EnvKey).(contract.Env)
	log.Println(envServ.AppEnv())
}
