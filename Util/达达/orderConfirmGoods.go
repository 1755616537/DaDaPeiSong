package 达达

import "github.com/gogf/gf/encoding/gjson"

//妥投异常之物品返回完成
//https://newopen.imdada.cn/#/development/file/abnormalConfirm?_k=liwas9
func (_data Config)OrderConfirmGoods(order_id string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//第三方订单编号
	data["order_id"]=order_id
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/order/confirm/goods")
}