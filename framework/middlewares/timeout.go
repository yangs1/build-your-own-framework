package middlewares

import (
	"context"
	"framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- err
				}
			}()

			ctx.Next()

			finish <- struct{}{} // 空结构体不占空间
		}()

		select {
		case err := <-panicChan:
			ctx.SetStatus(500).Json("err out")
			log.Println(err)
		case <-finish:
			log.Println("finish")
		case <-durationCtx.Done():
			log.Println("time out finish")

			ctx.SetStatus(500).Json("time out")
			ctx.SetHasTimeout()

			//ctx.GetResponse().Write([]byte("time out"))
			//ctx.Json(500, "time out")
		}

		return nil
	}
}
