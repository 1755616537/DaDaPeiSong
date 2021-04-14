package 达达

import "github.com/gogf/gf/encoding/gjson"

//新增订单
//https://newopen.imdada.cn/#/development/file/add?_k=y0d9ao
func (_data Config)OrderAddOrder(OrderType OrderType)(*gjson.Json,error)  {
	return _data.SetBody(gjson.New(OrderType)).SetSign().Post("/api/order/addOrder")
}
