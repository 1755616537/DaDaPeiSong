package 达达

import "github.com/gogf/gf/encoding/gjson"

//查询订单运费接口
//https://newopen.imdada.cn/#/development/file/readyAdd?_k=7ugnae
func (_data Config)OrderQueryDeliverFee(order_id string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//第三方订单编号
	data["order_id"]=order_id
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/order/queryDeliverFee")
}