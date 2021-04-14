package 达达

import "github.com/gogf/gf/encoding/gjson"

//取消订单
//https://newopen.imdada.cn/#/development/file/formalCancel?_k=2xy06z
func (_data Config)OrderFormalCancel(order_id string,cancel_reason_id int,cancel_reason string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//第三方订单编号
	data["order_id"]=order_id
	//取消原因ID
	data["cancel_reason_id"]=cancel_reason_id
	//取消原因(当取消原因ID为其他时，此字段必填)
	data["cancel_reason"]=cancel_reason
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/order/formalCancel")
}
