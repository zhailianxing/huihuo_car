package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
	"time"
	"os"
	"strconv"
	// "baseLayuiCMS2/models"
	"encoding/json"
	"baseLayuiCMS2/common"
)

type UploadController struct {
	beego.Controller
}

//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func create_local_file_path(FileName string) string{
	// TODO 加上 企业信息,去掉unix时间戳,不然很多文件夹
	file_path := "/tmp/" + strconv.Itoa(int(time.Now().Unix())) + "/"
	if !checkFileIsExist(file_path) {
		os.Mkdir(file_path, os.ModePerm)
	}
	return file_path + FileName
}

func (c *UploadController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
			logs.Error("getfile err ", err)
	}
	defer f.Close()
	local_path := create_local_file_path(h.Filename)
	logs.Info("UploadController  local_path:%v", local_path)
	c.SaveToFile("file", local_path) 
	// 不要删除.beego读取上传文件的二进制数据
	// data, err := ioutil.ReadAll(f)

	// 开发测试的时候就不再真正的上传图片
	// err, img_path := common.Upload(local_path)
	err, img_path := error(nil), "https://imimg.lilithcdn.com/avt/llc_avatar/test2.png"
	ret := Ret{}
	if err != nil {
		ret.Code = -1 
	}else{
		ret.Code = 0 
		ret.Data = img_path
	}
	retData, err := json.Marshal(ret)
	if err != nil {
		logs.Error("UploadController get 11111, err:", err)
	} else {
		logs.Info("UploadController  retData:%v", retData)
		c.Ctx.WriteString(string(retData))
	}
}

