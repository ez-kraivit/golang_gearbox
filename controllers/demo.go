package controllers

import (
	services "kgearboxs/services"
	"log"
	"os"

	"github.com/aymerick/raymond"
	"github.com/gogearbox/gearbox"
)

type demoController struct {
	demoService services.DemoService
}

func NewDemoController(s services.DemoService) demoController {
	return demoController{
		demoService: s,
	}
}

func (d demoController) Get(c gearbox.Context) {
	dat, err := os.ReadFile("sql/users.sql")
	if err != nil {
		log.Print(err)
	}
	ctx := map[string]string{
		"show":       "*",
		"table_name": "users",
	}
	result, err := raymond.Render(string(dat), ctx)
	if err != nil {
		log.Print("Please report a bug :)")
	}
	os.WriteFile("sql/users_demo.sql", []byte(result), 0644)
	c.SendJSON(result)
}

func (d demoController) Post(c gearbox.Context) {
	dat, err := os.ReadFile("sql/users.sql")
	if err != nil {
		log.Print(err)
	}
	ctx := map[string]string{
		"show":       "*",
		"table_name": "users",
	}
	result, err := raymond.Render(string(dat), ctx)
	if err != nil {
		log.Print("Please report a bug :)")
	}
	os.WriteFile("sql/users_demo.sql", []byte(result), 0644)
	c.SendJSON(result)
}
