package 达达

import "github.com/gogf/gf/encoding/gjson"

func messageConfirm()  {

}
//消息确认
//https://newopen.imdada.cn/#/development/file/merchantConfirm?_k=nl14e8
func (_data Config)MessageConfirm(messageType string,messageBody string)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	//消息类型（1：骑士取消订单推送消息）
	data["messageType"]=messageType
	//消息内容（json字符串）
	data["messageBody"]=messageBody
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/message/confirm")
}

//消息确认 骑士取消订单
//https://newopen.imdada.cn/#/development/file/transporterCancelOrder?_k=h5li6a
func (_data Config)MessageConfirm1(orderId string,dadaOrderId,isConfirm int)(*gjson.Json,error)  {
	var data map[string]interface{}=make(map[string]interface{})
	var messageBody map[string]interface{}=make(map[string]interface{})
	//商家第三方订单号
	messageBody["orderId"]=orderId
	//达达订单号
	messageBody["dadaOrderId"]=dadaOrderId
	//0:不同意，1:表示同意
	messageBody["isConfirm"]=isConfirm

	//消息类型（1：骑士取消订单推送消息）
	data["messageType"]=1
	//消息内容（json字符串）
	data["messageBody"]=messageBody
	return _data.SetBody(gjson.New(data)).SetSign().Post("/api/message/confirm")
}