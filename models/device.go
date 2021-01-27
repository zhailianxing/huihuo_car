package models

// import "github.com/astaxie/beego/orm"
import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/beego/beego/v2/core/logs"
)
//TODO:加索引
type Device struct {
	// 这里必须使用'json:"id"'. 这样json.Marsha就会将CompanyId变成company_id,在html中用的 company_id
	Id       int64  `json:"id"`
	CompanyId string `json:"company_id"`
	BrandId   string `json:"brand_id"`
	BrandType   string `json:"brand_type"`
	DeviceId string `json:"device_id"`
	UserId   int  `json:"user_id"`
	UserName   string  `json:"user_name"`
	UserPhone   string  `json:"user_phone"`
	Logo string  `json:"logo"`
	Ad string   `json:"ad"`
	IsDelete int  `json:"is_delete"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}


func DeviceGetByCompanyId(company_id string) (error, []*Device) {

	o := orm.NewOrm()

	obj := make([]*Device, 0)
	// ??? 为什么没有获取到数据
	_, err := o.QueryTable("device").Filter("company_id", company_id).All(&obj)
	logs.Error("models DeviceGet, company_id: %v, allDevices:%v", company_id, obj)

	if err != nil {
		return err, nil
	}

	return nil, obj
}

func DeviceGet(company_id, brand_id, device_id string) (error, *Device) {

	o := orm.NewOrm()

	obj := Device{}

	err := o.QueryTable("device").Filter("company_id", company_id).Filter("brand_id", brand_id).Filter("device_id", device_id).One(&obj)
	if err != nil {
		return err, nil
	}

	return nil, &obj
}

func UpdatePicture(company_id, brand_type, logo, ad string) error{
	o := orm.NewOrm()

	// d := &Device{}
	// d.CompanyId = company_id
	// d.BrandType = brand_type
	// // d.DeviceId = device_id
	// d.Ad = ad
	// d.Logo = logo
	// _, err := o.Update(d)  // 用update函数必须设置主键id字段

	_, err := o.QueryTable("device").Filter("company_id", company_id).Filter("brand_type", brand_type).Update(orm.Params{
		"ad": ad,
		"logo": logo,
	})
	return err
}

//
//// 修改
//func Update(city string, num int) error {
//
//	o := orm.NewOrm()
//
//	_, err := o.QueryTable("fei_yan").Filter("city", city).Update(orm.Params{
//		"num": num,
//	})
//	return err
//}
//
//// 插入
//func Add(city string, num int) error {
//
//	o := orm.NewOrm()
//
//	data := &FeiYan{City: city, Num: int64(num)}
//
//	_, err := o.Insert(data)
//
//	return err
//
//}
//
//// 删除
//func Del(id int) error {
//
//	o := orm.NewOrm()
//
//	//data := &FeiYan{Id: int64(id)}
//	//_, err := o.Delete(data)
//
//	data := FeiYan{Id: int64(id)}
//	_, err := o.Delete(&data)
//
//	return err
//
//}
