package framework

import (
	"context"
	"fmt"
	"framework/contract"
	"framework/provider/app"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestNewTree(t *testing.T) {
	var arr2 = []struct{}{}
	var arr1 = []struct{}{}
	t.Log(append(arr1, arr2...))
}

func TestGoroutine(t *testing.T) {
	now := time.Now()
	total := 10
	var wg sync.WaitGroup
	wg.Add(total)
	//整体业务超时时间为2s,如果2s内有某个服务超时,则统一退出
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < total; i++ {
		ctx2 := context.WithValue(ctx, "key", i)
		go func() {
			defer wg.Done()
			requestWork(ctx2)
		}()
	}
	wg.Wait()
	time.Sleep(8 * time.Second)
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}

func hardWork(t int) error {
	//每个业务执行时间不一致
	time.Sleep(time.Duration(t) * time.Second)
	return nil
}
func requestWork(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- hardWork(ctx.Value("key").(int))
	}()
	select {
	case <-done:
		fmt.Println("job done", ctx.Value("key"))
		return nil
	case <-ctx.Done():
		fmt.Println("ctx time up", ctx.Value("key"))
		return nil
	}
}

func TestProvider(t *testing.T) {
	container := NewLsfContainer()
	container.Bind(&app.LsfAppProvider{})

	lsfapp := container.MustMake(contract.AppKey).(contract.App)
	log.Println(lsfapp.BaseFolder())
}
