package main

import (
	"GinVueAdmin/server/core"
	"GinVueAdmin/server/global"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper()

	global.GVA_LOG = core.Zap()

	zap.ReplaceGlobals(global.GVA_LOG)
}
