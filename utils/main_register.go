package utils

import (
	"kgearboxs/routers"
	"kgearboxs/structs"

	"github.com/gogearbox/gearbox"
)

type MainRegister interface {
	DemoRouter() *gearbox.Route
}

type mainRegister struct {
	gb gearbox.Gearbox
}

func NewMainRegister(gb gearbox.Gearbox) mainRegister {
	return mainRegister{
		gb: gb,
	}
}

func (m mainRegister) DemoRouter() []*gearbox.Route {

	DemoRouter := routers.NewDemoRouter
	DemoData := structs.SRegisterRouter{
		Router: m.gb,
	}
	routers := m.gb.Group("/demo", DemoRouter(DemoData).Start())

	return routers
}
