package models

import (
	"baseLayuiCMS2/common"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/client/orm"
)

type Menu struct {
	Id    int64  `json:"id"`
	Pid   int    `json:"parentId"` // dtree要求，此字段名字必须是parentId
	Title string `json:"title"`
	Href  string `json:"href"`

	Spread    int    `json:"spread"`
	Target    string `json:"target"`
	Icon      string `json:"icon"`
	Available int    `json:"available"`
}

func GetMenuList(id int64) []*Menu {

	// TODO: 根据id属于哪种权限，然后拿到属于此权限下的菜单
	o := orm.NewOrm()

	m := make([]*Menu, 0)
	_, err := o.QueryTable("menu").Filter("available", common.AVAILABLE).All(&m)

	if err != nil {
		logs.Error("GetMenuList error,err:", err)
	}
	return m
}

func GetListByPage(parentId int, serarchMenuName string, page, limit int) (error, []*Menu) {

	o := orm.NewOrm()

	m := make([]*Menu, 0)
	offset := (page - 1) * limit

	var err error

	if parentId <= 0 {
		if len(serarchMenuName) <= 0 {
			_, err = o.QueryTable("menu").Filter("available", common.AVAILABLE).Offset(offset).Limit(limit).All(&m)
		} else {
			//  WHERE title LIKE '%...%'
			_, err = o.QueryTable("menu").Filter("available", common.AVAILABLE).Filter("title__icontains", serarchMenuName).Offset(offset).Limit(limit).All(&m)
		}

	} else {

		//orCond := orm.NewCondition()
		//orCond.And("id", parentId).Or("pid", parentId)
		//cond := orm.NewCondition()
		//cond.And("available", common.AVAILABLE).AndCond(orCond)
		//// where available = 1 and (id=? or parentId=?)
		//_, err = o.QueryTable("menu").SetCond(cond).Offset(offset).Limit(limit).All(&m)
		//var menus []*Menu
		menus := make([]*Menu, 0)
		var queryStr string
		if len(serarchMenuName) <= 0 {
			queryStr = fmt.Sprintf("SELECT *  FROM menu WHERE available = 1 and (id= %d or pid = %d )", parentId, parentId)
		} else {
			queryStr = fmt.Sprintf("SELECT *  FROM menu WHERE available = 1 and (id= %d or pid = %d ) and title like '%%%s%%' ", parentId, parentId, serarchMenuName)
		}

		fmt.Println("queryStr:", queryStr)
		//num, err := o.Raw("SELECT *  FROM menu WHERE available = 1 and (id= ? or pid = ? )", parentId, parentId).QueryRows(&menus)
		num, err := o.Raw(queryStr).QueryRows(&menus)
		if err != nil {
			fmt.Println("user nums: ", num)
			return err, menus
		}
		return nil, menus
	}

	if err != nil {
		logs.Error("GetListByPage error,err:", err)
		return err, m
	}
	return nil, m
}

func GetMenuCount(parentId int) int64 {

	o := orm.NewOrm()
	var err error
	var count int64
	if parentId <= 0 {
		count, err = o.QueryTable("menu").Count()
	} else {
		//res, err := o.Raw("SELECT count(*)  FROM menu WHERE available = 1 and (id= ? or pid = ? )", parentId, parentId).Exec()
		menus := make([]*Menu, 0)
		num, err := o.Raw("SELECT *  FROM menu WHERE available = 1 and (id= ? or pid = ? )", parentId, parentId).QueryRows(&menus)
		if err != nil {
			return 0
		}
		return num
	}

	if err != nil {
		logs.Error("GetListByPage error,err:", err)
		return 0
	}
	return count
}

func GetSecondMenu() (error, []*Menu) {

	o := orm.NewOrm()

	menus := make([]*Menu, 0)
	queryStr := fmt.Sprintf("SELECT *  FROM menu WHERE available = 1 and (pid= 0 or pid = 1 )")
	_, err := o.Raw(queryStr).QueryRows(&menus)
	if err != nil {
		return err, menus
	}
	return nil, menus
}

func AddMenu(m map[string]interface{}) error {

	o := orm.NewOrm()

	d := &Menu{}
	d.Available = m["available"].(int)
	d.Href = m["href"].(string)
	d.Icon = m["icon"].(string)
	d.Pid = m["pid"].(int)

	d.Spread = m["spread"].(int)
	d.Target = m["target"].(string)
	d.Title = m["title"].(string)

	i, _ := o.QueryTable("menu").PrepareInsert()
	_, err := i.Insert(d)

	if err != nil {
		return err
	}
	return nil
}

func EditMenu(m map[string]interface{}) error {
	o := orm.NewOrm()

	d := &Menu{}
	d.Id = int64(m["id"].(int))
	d.Available = m["available"].(int)
	d.Href = m["href"].(string)
	d.Icon = m["icon"].(string)
	d.Pid = m["pid"].(int)

	d.Spread = m["spread"].(int)
	d.Target = m["target"].(string)
	d.Title = m["title"].(string)

	_, err := o.Update(d)
	return err
}
