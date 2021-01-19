package routers

import (
	"baseLayuiCMS2/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
		// beego.Router("/", &controllers.MainController{})
		beego.Router("/test", &controllers.TestController{})

		beego.Router("/index", &controllers.IndexController{})
		beego.Router("/main", &controllers.MainController{})
	
		beego.Router("/login", &controllers.LoginController{})
	
		beego.Router("/layuiCmsluyun", &controllers.LayuiCmsluyunController{})
	
		beego.Router("/menuList", &controllers.MenuController{})
		// 只可以用 get/post/head等访问"127.0.0.1:8080/device"
		beego.Router("/device", &controllers.DeviceController{})
		// 自定义路径.支持 get/post/head等访问"127.0.0.1:8080/device/DeviceGetByCompanyId"
		beego.AutoRouter(&controllers.DeviceController{})

		beego.Router("/upload", &controllers.UploadController{})

}
