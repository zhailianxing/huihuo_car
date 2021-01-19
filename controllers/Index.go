package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	user := c.GetSession("userInfo")
	// logs("")
	logs.Info("user: %v", user)
	if user != nil { //兼容
		c.Data["user"] = user
		c.TplName = "index.html"
	}else {
		c.TplName = "login/login.html"
	}
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//  测试我的session
	//testUser := c.GetSession("userInfo")
	//fmt.Println("MainController  user:", testUser)

	c.TplName = "main.html"
}
