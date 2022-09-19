//go:build wireinject
// +build wireinject

package engine

import "github.com/google/wire"

func InitGinManager() *GinManager {
	panic(wire.Build(GinManagerSet))
}
