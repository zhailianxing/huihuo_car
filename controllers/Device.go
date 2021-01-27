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

// 列表页面
func (c *DeviceController) Get() {
	company := c.GetSession("userInfo")
	logs.Info("DeviceController DeviceList company: %v", company)
	if company != nil { //兼容
		c.TplName = "device/list.html"
	}else {
		c.TplName = "login/login.html"
	}
}
// 编辑页面
func (c *DeviceController) EditHtml() {
	c.TplName = "device/edit.html"
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
		logs.Error("DeviceController DeviceGetByCompanyId, err:", err)
	} else {
		logs.Info("DeviceController DeviceGetByCompanyId success, retData:%v", retData)
		c.Ctx.WriteString(string(retData))
	}
}

func (c *DeviceController) Post() {
	company := c.GetSession("userInfo")
	logs.Info("DeviceController Post company: %v", company)
	if company != nil { 
		company_id := c.GetString("company_id")
		brand_type := c.GetString("brand_type")
		// device_id := c.GetString("device_id")
		logo := c.GetString("logo")
		ad := c.GetString("ad")
		logs.Info("DeviceController Post, company_id:%s, brand_type:%s, logo:%s, ad:%s ", company_id, brand_type, logo, ad)
		err := models.UpdatePicture(company_id, brand_type, logo, ad)
		if err == nil {
			c.Ctx.WriteString(success())
		}else {
			logs.Error("DeviceController UpdatePicture error: %v", err)
			c.Ctx.WriteString(error(-1, err.Error()))
		}
	}else {
		c.TplName = "login/login.html"
	}
}


