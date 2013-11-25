package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/spartacusX/ambuild/controllers"
)

func main() {
	beego.Router("/index", &controllers.MainController{})
	beego.Run()
}
