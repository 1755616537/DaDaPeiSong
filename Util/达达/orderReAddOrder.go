package 达达

import "github.com/gogf/gf/encoding/gjson"

//重新发布订单
//https://newopen.imdada.cn/#/development/file/reAdd?_k=e3wmsj
func (_data Config)OrderReAddOrder(OrderType OrderType)(*gjson.Json,error)  {
	return _data.SetBody(gjson.New(OrderType)).SetSign().Post("/api/order/reAddOrder")
}