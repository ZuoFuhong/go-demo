package common

import "sync"

// session管理

var manager sync.Map

func store(userId string, ctx *ConnContext) {
	manager.Store(userId, ctx)
}

func load(userId string) *ConnContext {
	value, ok := manager.Load(userId)
	if ok {
		return value.(*ConnContext)
	}
	return nil
}

func delete(userId string) {
	manager.Delete(userId)
}
