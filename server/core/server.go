package core

import "GinVueAdmin/server/initialize"

func RunServer() {
	router := initialize.Routers()
	_ = router.Run()
}
