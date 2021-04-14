package 达达

import "github.com/gogf/gf/encoding/gjson"

//查询订单运费接口
//https://newopen.imdada.cn/#/development/file/readyAdd?_k=7ugnae
func (_data Config)OrderQueryDeliverFee(OrderType OrderType)(*gjson.Json,error)  {
	return _data.SetBody(gjson.New(OrderType)).SetSign().Post("/api/order/queryDeliverFee")
}