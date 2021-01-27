
package controllers
import (
	"encoding/json"
)

// 数据表格(table)需要的数据返回格式
type Ret struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`

	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

func success() string{
	ret := Ret{}
	ret.Code = 0
	ret.Message = "success"
	retData, _ := json.Marshal(ret)
	return string(retData)
}

func error(err_code int, err_msg string) string{
	ret := Ret{}
	ret.Code = err_code
	ret.Message = err_msg
	retData, _ := json.Marshal(ret)
	return string(retData)
}
