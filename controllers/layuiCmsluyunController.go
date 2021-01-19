package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)
type LayuiCmsluyunController struct {
	beego.Controller
}

func (c *LayuiCmsluyunController) Get() {

	c.TplName = "layuiCmsluyun/index.html"
}
