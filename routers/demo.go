package routers

import (
	"kgearboxs/controllers"
	"kgearboxs/middleware"
	"kgearboxs/repository"
	"kgearboxs/services"
	structs "kgearboxs/structs"

	"github.com/gogearbox/gearbox"
	"gorm.io/gorm"
)

type DemoRouter interface {
	Start(data demoRouter) gearbox.Gearbox
}

type demoRouter struct {
	router gearbox.Gearbox
	db     *gorm.DB
}

func NewDemoRouter(data structs.SRegisterRouter) demoRouter {
	return demoRouter{
		router: data.Router,
	}
}

func (d demoRouter) Start() []*gearbox.Route {

	usersRepository := repository.NewUsersRepository(d.db)
	demoService := services.NewDemoService(usersRepository)
	demoController := controllers.NewDemoController(demoService)

	return []*gearbox.Route{
		d.router.Get("", demoController.Get),
		d.router.Post("", middleware.AuthorizedMiddleware, demoController.Post),
	}
}
