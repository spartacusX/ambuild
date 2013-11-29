package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"os/exec"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

type MainController struct {
	beego.Controller
}

type BuildParameters struct {
	Platforms []string
	Encodings []string
	Languages []string
	Sign      string
	IncBuild  string
	MvBuild   string
	EnableMOE string
	EnableCSA string
}

const (
	BUILD_READY = "active"
	BUILD_BUSY  = "disabled"
)

var BUILD_STATUS = BUILD_READY

func (this *MainController) Get() {
	fmt.Println("Request Method: Get")
	this.Data["Language"] = this.GetSession("Language")
	this.Data["Encoding"] = this.GetSession("Encoding")
	this.Data["Revision"] = this.GetSession("Revision")

	if this.GetSession("BuildLaunched") == "success" {
		this.Data["showhidden"] = "show"
	} else {
		this.Data["showhidden"] = "hidden"
	}

	this.Data["BuildStatus"] = BUILD_STATUS
	this.Layout = "layout.html"
	this.TplNames = "build.tpl"
}

func (this *MainController) Post() {
	fmt.Println("Request Method: Post")

	this.Ctx.Request.ParseForm()
	var bp BuildParameters
	bp.Platforms = this.GetStrings("PlatForm")
	bp.Encodings = this.GetStrings("Encoding")
	bp.Languages = this.GetStrings("Language")
	bp.Sign = this.GetString("Sign")
	bp.IncBuild = this.GetString("IncBuild")
	bp.MvBuild = this.GetString("MVBuild")
	bp.EnableMOE = this.GetString("EnableMOE")
	bp.EnableCSA = this.GetString("EnableCSA")

	//revision := this.Ctx.Request.Form.Get("BuildRevision")

	fmt.Println(bp)
	ParameterValidation(bp)
	// this.SetSession("Language", language)
	// this.SetSession("Encoding", encode)
	// this.SetSession("Revision", revision)

	//beego.Info(this.Ctx.Request.Form)

	err := LaunchBuild(bp)
	if err == nil {
		this.SetSession("BuildLaunched", "success")
		BUILD_STATUS = BUILD_BUSY
	}

	go BuildProgressCheck(bp.Platforms[0], bp.Encodings[0])

	this.Ctx.Redirect(302, "/index")
}

func LaunchBuild(bp BuildParameters) error {
	// cmd := exec.Command("bash", "/home/build/build.sh", language, encode, revision)
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("The build has been launched successfully.")
	// }
	// go EmailNotify()

	return nil
}

func EmailNotify() error {
	user := "libin.tian@hp.com"
	password := "xxxxx"
	host := "casarray1.austin.hp.com:465"
	to := "libin_tian@163.com;libin.tian822@gmail.com"

	subject := "Test send email by golang"

	body := `
 	    <html>
 	    <body>
 	    <h3>
 	    "Test send email by golang"
 	    </h3>
 	    </body>
 	    </html>
 	    `
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

	return err
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func BuildProgressCheck(platform, encoding string) {
	const SUCCESS = []byte("009900")
	const FAILED = []byte("FF0000")
	const target = []byte(`FONT COLOR="#`)
	url := "http://amdbserver.chn.hp.com/acall/build.php"
	//target := []byte("http://amdbserver.chn.hp.com/AssetManager/main/9.40.")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Retrieve build log failed: ", err)
	} else {
		content, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			index := bytes.Index(content, target)

			buildstatus := content[index+len(target) : index+len(target)+6]
			if bytes.Compare(buildstatus, SUCCESS) == 0 {
				fmt.Println("Build finished successfully.")
				BUILD_STATUS = BUILD_READY
			} else if bytes.Compare(buildstatus, FAILED) == 0 {
				fmt.Println("Build finished but failed.")
				BUILD_STATUS = BUILD_READY
			} else {
				BUILD_STATUS = BUILD_BUSY
			}
			fmt.Println(bytes.NewBuffer(buildstatus).String())
		}
	}

}
