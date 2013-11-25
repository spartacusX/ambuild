package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println("Request Method: Get")

	this.Layout = "layout.html"
	this.TplNames = "build.tpl"
}

func (this *MainController) Post() {
	fmt.Println("Request Method: Post")
	//数据处理
	this.Ctx.Request.ParseForm()
	language := this.Ctx.Request.Form.Get("BuildLanguage")
	encode := this.Ctx.Request.Form.Get("BuildType")
	revision := this.Ctx.Request.Form.Get("BuildRevision")

	fmt.Println(language)
	fmt.Println(encode)
	fmt.Println(revision)

	//beego.Info(this.Ctx.Request.Form)

	LaunchBuild(language, encode, revision)

	this.Ctx.Redirect(302, "/index")
}

func LaunchBuild(language string, encode string, revision string) {
	out, err := exec.Command("cmd", "ver").Output()
	//err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	bbuf := bytes.NewBuffer(out)
	fmt.Println("Done.", bbuf.String())
}
