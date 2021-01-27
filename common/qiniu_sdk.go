package common

import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = "GQBMiT0Wy63GIEFUzh_-93IXvds3v9fcowxcWBPo" // 七牛的accessKey 去七牛后台获取
	secretKey = "3aQdDQ2xENbdn1P382bdZC5t3dTRSGViABqGfNOF" // 七牛的secretKey 去七牛后台获取
	bucket    = "im-img"                                // 上传空间 去七牛后台创建
	host = "https://imimg.lilithcdn.com/"
	// 需要上传的文件
	// localFile := "/Users/lilithgames/Downloads/icon_lilith_avatar@3x.png"
)

// 上传本地文件. localFile指定路径
func Upload(local_file_path string) (error, string){
	// 鉴权
	mac := qbox.NewMac(accessKey, secretKey)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	//上传文件失败,原因: incorrect region, please use up-z0.qiniup.com
	// 这里要改为"ZoneHuadong". 参考:https://developer.qiniu.com/kodo/1671/region-endpoint-fq
	cfg.Zone = &storage.ZoneHuadong //指定上传的区域
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      //是否使用CDN上传加速

	// 七牛key
	// qiniuKey := "avt/llc_avatar/test2.png"
	qiniuKey := local_file_path //qiniuKey 只是用来表示上传后url的链接. 例:上传后的图片:"https://imimg.lilithcdn.com/avt/llc_avatar/test2.png"

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, local_file_path, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		return err, ""
	}
	fmt.Println("上传成功,key为:", ret.Key, ret.Hash)
	return nil, host + ret.Key
}

// TODO:使用二进制图片上传
// func UploadByFileReader(File io.ReaderAt) (error, string){
// TODO
// }