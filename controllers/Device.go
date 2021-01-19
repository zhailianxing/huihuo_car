package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
	"baseLayuiCMS2/models"
	"encoding/json"
)

type DeviceController struct {
	beego.Controller
}

// 页面
func (c *DeviceController) Get() {
	company := c.GetSession("userInfo")
	logs.Info("DeviceController DeviceList company: %v", company)
	if company != nil { //兼容
		c.TplName = "device/list.html"
	}else {
		c.TplName = "login/login.html"
	}
}
// 接口
func (c *DeviceController) DeviceGetByCompanyId() {
	company := c.GetSession("userInfo")
	logs.Info("DeviceController DeviceGetByCompanyId company: %v", company)
	ret := Ret{}
	if company != nil { //兼容
		company_id := company.(*models.Company).CompanyId
		_, allDevices := models.DeviceGetByCompanyId(company_id)
		ret.Code = 0
		ret.Message = "success"
		ret.Count = int64(len(allDevices))
		logs.Error("DeviceController, company_id: %v, allDevices:%v", company_id, allDevices)
		ret.Data = allDevices
	}else{
		ret.Code = 1
		ret.Message = "error"
	}
	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("DeviceController get 11111, err:", err)
	} else {
		logs.Error("DeviceController get 222222, retData:%v", retData)
		c.Ctx.WriteString(string(retData))
	}
}


