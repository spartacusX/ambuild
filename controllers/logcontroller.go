package controllers

import (
	//"bytes"
	"fmt"
	"github.com/astaxie/beego"
	//"html/template"
	//"io/ioutil"
)

type LogController struct {
	beego.Controller
}

func (this *LogController) Get() {
	fmt.Println("Request method: Get in LogController.")

	fmt.Println(this.Ctx.Request.RequestURI)
	this.Layout = "log/prog.html"
	this.TplNames = "build.tpl"
	// logPath := "/opt/fdcmfiler/lastbuild/AssetManager/main/log/prog.html"

	// logByte, err := ioutil.ReadFile(logPath)
	// logOutput := "Build log"
	// if err != nil {
	// 	logOutput = "Reading log failed: " + err.Error()
	// 	fmt.Println(logOutput)
	// } else {
	// 	buf := bytes.NewBuffer(logByte)
	// 	logOutput = buf.String()
	// }
	// this.Data["xml"] = template.HTML(logOutput)
	// this.ServeXml()
}
