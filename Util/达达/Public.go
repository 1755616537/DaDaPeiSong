package 达达

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"strconv"
	"strings"
	"time"
	"达达配送/Util"
)

//https://newopen.imdada.cn/#/quickStart/develop/mustRead?_k=lzfb26
type Config struct {
	//业务参数，JSON字符串
	Body string
	//请求格式，暂时只支持json
	Format string
	//时间戳,单位秒，即unix-timestamp
	Timestamp string
	//签名Hash值，参见：接口签名规则
	Signature string
	//应用Key，对应开发者账号中的app_key
	App_key string
	//API版本
	V string
	//商户编号（创建商户账号分配的编号）	测试环境默认为：73753
	Source_id string

	App_secret string

	//是否使用测试环境
	IsTest bool
}

//达达 初始化配置信息 Source_id,测试环境默认为=73753
func GetDaDaPeiSongRun(App_key,App_secret,Source_id string)Config  {
	return Config{
		Format: "json",
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		App_key: App_key,
		V: "1.0",
		Source_id: Source_id,
		App_secret: App_secret,

		IsTest: true,
	}
}

//达达 获取签名
//https://newopen.imdada.cn/#/quickStart/develop/safety?_k=tazotb
func (_data Config)SetSign()Config  {
	sign:=""
	sign=fmt.Sprint(sign,
		"app_key",_data.App_key,
		"body",_data.Body,
		"format",_data.Format,
		"source_id",_data.Source_id,
		"timestamp",_data.Timestamp,
		"v",_data.V,
		)
	sign=fmt.Sprint(_data.App_secret,sign,_data.App_secret)
	sign=Util.GetMD5(sign)
	sign=strings.ToUpper(sign)
	_data.Signature=sign

	return _data
}

func (_data Config)SetTest(is bool)Config  {
	_data.IsTest=is
	return _data
}

func (_data Config)SetBody(Body *gjson.Json)Config  {
	_data.Body=Body.MustToJsonString()
	return _data
}

func (_data Config)Post(urll string)(*gjson.Json,error)  {
	//正式环境
	url:="https://newopen.imdada.cn"
	if _data.IsTest {
		//测试环境
		url="https://newopen.qa.imdada.cn"
	}

	type Config_ struct {
		//业务参数，JSON字符串
		Body string `json:"body"`
		//请求格式，暂时只支持json
		Format string `json:"format"`
		//时间戳,单位秒，即unix-timestamp
		Timestamp string `json:"timestamp"`
		//签名Hash值，参见：接口签名规则
		Signature string `json:"signature"`
		//应用Key，对应开发者账号中的app_key
		App_key string `json:"app_key"`
		//API版本
		V string `json:"v"`
		//商户编号（创建商户账号分配的编号）	测试环境默认为：73753
		Source_id string `json:"source_id"`
	}
	var data Config_
	data.Body=_data.Body
	data.Format=_data.Format
	data.Timestamp=_data.Timestamp
	data.Signature=_data.Signature
	data.App_key=_data.App_key
	data.V=_data.V
	data.Source_id=_data.Source_id

	_, _Ret, err := Util.HTTPGet2("POST", fmt.Sprint(url, urll), data)
	if err != nil {
		return nil,errors.New("请求错误")
	}

	_RetJson:=gjson.New(_Ret)
	if _RetJson.GetString("status")!="success" {
		code:=_RetJson.GetString("code")
		msg:=_RetJson.GetString("msg")
		if code=="" || msg=="" {
			return _RetJson,errors.New("请求错误")
		}
		return _RetJson,errors.New(fmt.Sprint("请求错误.",code,msg))
	}

	return _RetJson,nil
}

//订单 数据类型
type OrderType struct {
	//门店编号，门店创建后可在门店列表和单页查看
	Shop_no string `json:"shop_no"`
	//第三方订单ID
	Origin_id string `json:"origin_id"`
	//订单所在城市的code
	City_code string `json:"city_code"`
	//订单金额（单位：元）
	Cargo_price float64 `json:"cargo_price"`
	//是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	Is_prepay int `json:"is_prepay"`
	//收货人姓名
	Receiver_name string `json:"receiver_name"`
	//收货人地址
	Receiver_address string `json:"receiver_address"`
	//收货人地址纬度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度，高德地图坐标拾取器）
	Receiver_lat float64 `json:"receiver_lat"`
	//收货人地址经度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度，高德地图坐标拾取器)
	Receiver_lng float64 `json:"receiver_lng"`
	//回调URL
	Rallback string `json:"callback"`
	//订单重量（单位：Kg）
	Cargo_weight float64 `json:"cargo_weight"`
	//收货人手机号（手机号和座机号必填一项）
	Receiver_phone string `json:"receiver_phone"`
	//收货人座机号（手机号和座机号必填一项）
	Receiver_tel string `json:"receiver_tel"`

	//小费（单位：元，精确小数点后一位）
	Tips float64 `json:"tips"`
	//订单备注
	Info string `json:"info"`
	//订单商品类型：食品小吃-1,饮料-2,鲜花-3,文印票务-8,便利店-9,水果生鲜-13,同城电商-19, 医药-20,蛋糕-21,酒品-24,小商品市场-25,服装-26,汽修零配-27,数码-28,小龙虾-29,火锅-51,其他-5
	Cargo_type int `json:"cargo_type"`
	//订单商品数量
	Cargo_num int `json:"cargo_num"`
	//发票抬头
	Invoice_title string `json:"invoice_title"`
	//订单来源标示（只支持字母，最大长度为10）
	Origin_mark string `json:"origin_mark"`
	//订单来源编号，最大长度为30，该字段可以显示在骑士APP订单详情页面，示例：origin_mark_no:"#京东到家#1"达达骑士APP看到的是：#京东到家#1
	Origin_mark_no string `json:"origin_mark_no"`
	//是否使用保价费（0：不使用保价，1：使用保价； 同时，请确保填写了订单金额（cargo_price））商品保价费(当商品出现损坏，可获取一定金额的赔付)保费=配送物品实际价值*费率（5‰），配送物品价值及最高赔付不超过10000元， 最高保费为50元（物品价格最小单位为100元，不足100元部分按100元认定，保价费向上取整数， 如：物品声明价值为201元，保价费为300元*5‰=1.5元，取整数为2元。）若您选择不保价，若物品出现丢失或损毁，最高可获得平台30元优惠券。 （优惠券直接存入用户账户中）。
	Is_use_insurance int `json:"is_use_insurance"`
	//收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
	Is_finish_code_needed string `json:"is_finish_code_needed"`
	//预约发单时间（预约时间unix时间戳(10位),精确到分;整分钟为间隔，并且需要至少提前5分钟预约，可以支持未来3天内的订单发预约单。）
	Delay_publish_time int `json:"delay_publish_time"`
	//是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑士只能配送此订单至完成，同时，也会相应的增加配送费用）
	Is_direct_delivery string `json:"is_direct_delivery"`
	//订单商品明细
	Product_list []ProductListType `json:"product_list"`
	//货架信息,该字段可在骑士APP订单备注中展示
	Pick_up_pos string `json:"pick_up_pos"`
}

type ProductListType struct {
	//商品名称，限制长度128
	Sku_name string `json:"sku_name"`
	//商品编码，限制长度64
	Src_product_no string `json:"src_product_no"`
	//商品数量，精确到小数点后两位
	Count float64 `json:"count"`
	//商品单位，默认：件
	Unit string `json:"unit"`
}

//达达配送 回调 数据类型
type SetBodyType struct {
	//返回达达运单号，默认为空
	Client_id string `json:"client_id"`
	//添加订单接口中的origin_id值
	Order_id string `json:"order_id"`
	//订单状态(待接单＝1,待取货＝2,配送中＝3,已完成＝4,已取消＝5, 指派单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 骑士到店=100,创建达达运单失败=1000 可参考文末的状态说明）
	Order_status int `json:"order_status"`
	//订单取消原因,其他状态下默认值为空字符串
	Cancel_reason string `json:"cancel_reason"`
	//订单取消原因来源(1:达达配送员取消；2:商家主动取消；3:系统或客服取消；0:默认值)
	Cancel_from int `json:"cancel_from"`
	//更新时间，时间戳除了创建达达运单失败=1000的精确毫秒，其他时间戳精确到秒
	Update_time int64 `json:"update_time"`
	//对client_id, order_id, update_time的值进行字符串升序排列，再连接字符串，取md5值
	Signature string `json:"signature"`

	//达达配送员id，接单以后会传
	Dm_id int `json:"dm_id"`
	//配送员姓名，接单以后会传
	Dm_name string `json:"dm_name"`
	//配送员手机号，接单以后会传
	Dm_mobile string `json:"dm_mobile"`
	//收货码
	Finish_code string `json:"finish_code"`
}

//回调
//https://newopen.imdada.cn/#/development/file/order?_k=55zbmn
func (_data Config)Set(Body string)(*SetBodyType,error)  {
	var data SetBodyType
	err := json.Unmarshal([]byte(Body), &data)
	if err != nil {
		return nil,err
	}
	return &data,nil
}