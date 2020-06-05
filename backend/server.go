package main

import (
	"gousers/core"
	"gousers/enums"
	"gousers/logger"
	pingmodule "gousers/modules/ping"
	usermodule "gousers/modules/user"
)

var (
	port      = enums.GetPort()
	apiPrefix = enums.GetAPIPrefix()
)

func main() {
	logger.Scaffold()
	router := core.Bootstrap(apiPrefix)
	bootstrap(router)
	router.Run(port)
}

// add your routes here
func bootstrap(api *core.GousersRouter) {
	// Bootstrap all the Models
	pingmodule.BootstrapRoutes(api)
	usermodule.BootstrapRoutes(api)
}
