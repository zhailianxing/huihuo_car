package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}

// 页面
func (c *TestController) Get() {
	c.TplName = "test/upload.html"
}



