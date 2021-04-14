package Test

import (
	"fmt"
	"testing"
	"达达配送/Util/达达"
)

func Test增单(t *testing.T) {
	_Ret, err := 达达.GetDaDaPeiSongRun("", "", "").OrderAddOrder(达达.OrderType{

	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//配送距离
	distance:=_Ret.GetFloat64("result.distance")
	//实际运费
	fee:=_Ret.GetFloat64("result.fee")
	fmt.Println(distance,fee)
}
