package controllers

// import (
// 	"beego-layui/common"
// 	"beego-layui/models"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/logs"
// )

import (
	"baseLayuiCMS2/models"
	"baseLayuiCMS2/common"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
)

type MenuController struct {
	beego.Controller
}

type MenuLink struct {
	Id    int64  `json:"id"`
	Pid   int    `json:"parentId"`
	Title string `json:"title"`
	Href  string `json:"href"`

	Spread int    `json:"spread"`
	Target string `json:"target"`
	Icon   string `json:"icon"`

	Children []*MenuLink `json:"children"`
}

func (c *MenuController) Get() {

	user := c.GetSession("userInfo")

	userStruct, _ := user.(models.User)
	list := models.GetMenuList(userStruct.Id)

	// 组装格式：
	ret := make([]*MenuLink, 0)
	for _, v := range list {
		if v.Pid == common.TOPPARENETID { // pid=1表示是 父菜单。其他的都是子菜单。 注意：我们这里只有二级菜单。如果多级，需要递归
			item := MenuLink{}
			item.Pid = v.Pid
			item.Id = v.Id
			item.Title = v.Title
			item.Href = v.Href
			item.Icon = v.Icon
			item.Spread = v.Spread
			item.Target = v.Target

			ret = append(ret, &item)

			for _, node := range list {
				if node.Pid == int(v.Id) {
					fmt.Println("pid:", node.Pid, ",Id:", v.Id)
					child := MenuLink{}
					child.Pid = node.Pid
					child.Id = node.Id
					child.Title = node.Title
					child.Href = node.Href
					child.Icon = node.Icon
					child.Spread = node.Spread
					child.Target = node.Target

					item.Children = append(item.Children, &child)
				}
			}
		}
	}

	//fmt.Println("over:")
	//for _, v := range ret {
	//	fmt.Println(v)
	//	for _, child := range v.Children {
	//		fmt.Println(child)
	//	}
	//}

	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuController get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}
}

// dtree需要的数据表格格式
type DtreeStatus struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
type Dtree struct {
	DtreeStatus `json:"status"`
	Data        interface{} `json:"data"`
}

func (c *MenuController) MangerDtreeList() {

	list := models.GetMenuList(-1)

	// 组装格式：
	data := make([]*models.Menu, 0)
	for _, v := range list {
		item := models.Menu{}
		item.Pid = v.Pid
		item.Id = v.Id
		item.Title = v.Title
		item.Href = v.Href
		item.Icon = v.Icon
		item.Spread = v.Spread
		item.Target = v.Target

		data = append(data, &item)
	}

	ret := Dtree{DtreeStatus: DtreeStatus{Code: 200, Message: "success"}, Data: data}

	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuMangerDtreeList get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}

}

// "菜单管理"
func (c *MenuController) Manger() {
	c.TplName = "menu/menuManger.html"
}

// 请求时必须要小写，即：/menu/mangerleft
func (c *MenuController) MangerLeft() {
	c.TplName = "menu/menuMangerLeft.html"
}

func (c *MenuController) MangerRight() {
	c.TplName = "menu/menuMangerRight.html"
}

// // 数据表格(table)需要的数据返回格式
// type Ret struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"msg"`

// 	Count int64       `json:"count"`
// 	Data  interface{} `json:"data"`
// }

// 数据表格加载所有菜单内容
func (c *MenuController) LoadAll() {

	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	menu_id, _ := c.GetInt("menu_id")
	serarchMenuName := c.GetString("serarchMenuName")

	err, data := models.GetListByPage(menu_id, serarchMenuName, page, limit)
	ret := Ret{}
	if err != nil {
		ret.Code = 1
		ret.Message = "error"
	} else {
		ret.Code = 0
		ret.Message = "success"
		ret.Count = models.GetMenuCount(menu_id)
		ret.Data = data
	}

	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuMangerDtreeList get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}
}

func (c *MenuController) GetSecondMenu() {

	err, data := models.GetSecondMenu()
	ret := Ret{}
	if err != nil {
		ret.Code = 1
		ret.Message = "error"
	} else {
		ret.Code = 0
		ret.Message = "success"
		ret.Data = data
	}

	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuMangerDtreeList get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}
}

func (c *MenuController) Add() {

	//  map[available:[0] href:[222] icon:[444] pid:[1] spread:[1] target:[333] title:[111]]
	available, _ := c.GetInt("available")
	href := c.GetString("href")
	icon := c.GetString("icon")
	pid, _ := c.GetInt("pid")
	spread, _ := c.GetInt("spread")
	target := c.GetString("target")
	title := c.GetString("title")

	m := map[string]interface{}{}
	m["available"] = available
	m["href"] = href
	m["icon"] = icon
	m["pid"] = pid
	m["spread"] = spread
	m["target"] = target
	m["title"] = title
	err := models.AddMenu(m)

	ret := Ret{Code: 0, Message: "success"}
	if err != nil {
		ret.Code = 1
		ret.Message = err.Error()
	}
	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuMangerDtreeList get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}

}

func (c *MenuController) Edit() {

	//  map[available:[0] href:[222] icon:[444] pid:[1] spread:[1] target:[333] title:[111]]
	id, _ := c.GetInt("id")

	available, _ := c.GetInt("available")
	href := c.GetString("href")
	icon := c.GetString("icon")
	pid, _ := c.GetInt("pid")
	spread, _ := c.GetInt("spread")
	target := c.GetString("target")
	title := c.GetString("title")

	m := map[string]interface{}{}
	m["id"] = id
	m["available"] = available
	m["href"] = href
	m["icon"] = icon
	m["pid"] = pid
	m["spread"] = spread
	m["target"] = target
	m["title"] = title
	err := models.EditMenu(m)

	ret := Ret{Code: 0, Message: "success"}
	if err != nil {
		ret.Code = 1
		ret.Message = err.Error()
	}
	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("MenuMangerDtreeList get, err:", err)
	} else {
		c.Ctx.WriteString(string(retData))
	}
}
