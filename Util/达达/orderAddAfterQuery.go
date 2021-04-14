package 达达

import "github.com/gogf/gf/encoding/gjson"

//查询运费后发单接口
//https://newopen.imdada.cn/#/development/file/readyAdd?_k=7ugnae
func (_data Config)orderAddAfterQuery(deliveryNo string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//平台订单编号
	data["deliveryNo"]=deliveryNo
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/order/addAfterQuery")
}
