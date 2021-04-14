package 达达

import "github.com/gogf/gf/encoding/gjson"

//商家投诉达达
//https://newopen.imdada.cn/#/development/file/complaintDada?_k=83wwba
func (_data Config)ComplaintDada(order_id string,reason_id int)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//第三方订单编号
	data["order_id"]=order_id
	//投诉原因ID
	data["reason_id"]=reason_id
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/complaint/dada")
}