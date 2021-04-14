package 达达

import "github.com/gogf/gf/encoding/gjson"

//订单详情查询
//https://newopen.imdada.cn/#/development/file/statusQuery?_k=215zvt
func (_data Config)OrderStatusQuery(OrderType OrderType)(*gjson.Json,error)  {
	return _data.SetBody(gjson.New(OrderType)).SetSign().Post("/api/order/status/query")
}
