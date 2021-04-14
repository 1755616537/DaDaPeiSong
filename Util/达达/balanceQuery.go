package 达达

import "github.com/gogf/gf/encoding/gjson"

//查询账户余额
//https://newopen.imdada.cn/#/development/file/balanceQuery?_k=a1n79r
func (_data Config)BalanceQuery(category int)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//查询运费账户类型（1：运费账户；2：红包账户，3：所有），默认查询运费账户余额
	data["category"]=category
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/balance/query")
}