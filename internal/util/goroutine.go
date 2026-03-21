package util

import (
	"context"
	"god-help-service/internal/util/logger"
	"runtime/debug"
)

// SafeGo 安全地启动协程，捕获并处理panic
func SafeGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 记录panic信息和堆栈
				stack := string(debug.Stack())
				logger.Errorf("协程panic:%s 堆栈:%s", err, stack)
			}
		}()
		f()
	}()
}

// SafeGoWithContext 带上下文的安全协程启动
func SafeGoWithContext(ctx context.Context, f func(context.Context)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 记录panic信息和堆栈
				stack := string(debug.Stack())
				logger.Errorf("协程panic:%s 堆栈:%s", err, stack)
			}
		}()
		f(ctx)
	}()
}

// SafeGoWithName 带名称的安全协程启动，便于日志识别
func SafeGoWithName(name string, f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 记录panic信息和堆栈
				stack := string(debug.Stack())
				logger.Errorf("协程[%s]panic:%s", name, stack)
			}
		}()
		f()
	}()
}
