package app

import (
	"framework"
	"framework/contract"
	"log"
	"testing"
)

func TestProvider(t *testing.T) {
	container := framework.NewLsfContainer()
	container.Bind(&LsfAppProvider{})

	lsfapp := container.MustMake(contract.AppKey).(contract.App)
	log.Println(lsfapp.BaseFolder())
}
