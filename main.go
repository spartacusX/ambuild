package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/spartacusX/ambuild/controllers"
)

func main() {
	beego.SessionOn = true

	beego.SetStaticPath("/views/log", "log")

	beego.Router("/index", &controllers.MainController{})
	beego.Router("/log/*", &controllers.LogController{})
	beego.Run()
}
