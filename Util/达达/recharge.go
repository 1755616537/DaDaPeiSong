package 达达

import "github.com/gogf/gf/encoding/gjson"

//获取充值链接
//https://newopen.imdada.cn/#/development/file/recharge?_k=ylsi2j
func (_data Config)Recharge(amount float64,category string,notify_url string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//充值金额（单位元，可以精确到分）
	data["amount"]=amount
	//生成链接适应场景（category有二种类型值：PC、H5）
	data["category"]=category
	//支付成功后跳转的页面（支付宝在支付成功后可以跳转到某个指定的页面，微信支付不支持）
	data["notify_url"]=notify_url
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/recharge")
}
