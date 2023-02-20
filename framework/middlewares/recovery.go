package middlewares

import "framework"

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatus(500).Json(err)
			}
		}()
		// 使用next执行具体的业务逻辑
		ctx.Next()

		return nil
	}
}
