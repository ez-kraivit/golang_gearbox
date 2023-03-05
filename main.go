package main

import (
	middleware "kgearboxs/middleware"
	utils "kgearboxs/utils"

	"github.com/gogearbox/gearbox"
)

func main() {

	// Setup gearbox
	gb := gearbox.New()

	corsMiddleware := middleware.CorsMiddleware
	logMiddleware := middleware.LogMiddleware

	// Register Plugins
	gb.Use(corsMiddleware)
	gb.Use(logMiddleware)

	// Static
	gb.Static("/static", "./public")

	mainRegister := utils.NewMainRegister(gb)
	mainRegister.DemoRouter()

	// Start service
	gb.Start("0.0.0.0:8080")
}
