package controllers

// import (
// 	"beego-layui/models"

// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/logs"
// )

import (
  "baseLayuiCMS2/models"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.TplName = "login/login.html"
}

func (p *LoginController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

func (c *LoginController) Post() {
	// name := c.Input().Get("userName")
	// pwd := c.Input().Get("password")

	name:= c.GetString("userName")
	pwd := c.GetString("password")

	logs.Info("name:", name, ",pwd:", pwd)

	err, user := models.GetCompany(name, pwd)
	if err != nil {
		// 版本1：直接重定向，没有提示信息。
		//c.Redirect("/login", 302)
		// 版本2：先搞个提示信息，再重定向
		c.History("登陆失败", "")
		// TODO 版本3： 登陆页面发送ajax获取接口返回的json数据。 在页面内处理
	} else {
		c.SetSession("userInfo", user)
		c.Redirect("/index", 301)

	}
}
