
package controllers


// 数据表格(table)需要的数据返回格式
type Ret struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`

	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}