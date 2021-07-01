package main

import (
	"fmt"
	"math"
	"testing"
)

const str ="Go101.org"
var a byte =1<<len(str)/128
var b byte =1<<len(str[:])/128
func Test_byte(t *testing.T){
	//移位运算在go中有个特殊的规定
	//1 移位运算的右侧必须是一个整数类型,或者是一个unit类型表示的无类型的常量(其实也就是字符串array之类的len)
	//2 如果位运算的左侧是一个 无类型的常亮/一个常亮那么结果是整数类型
	//3 如果位运算的左侧是一个 无类型的常亮/一个非常亮表达式 那么结果会是一个byte类型.1被隐士的转换为byte类型了
	//4 1<<(str[:]) 1被转化为byte类型 超出了最大值结果为0/128
	//5 感觉虽然是这样但是go的编译器还是不够聪明啊,明显是可以报错的,已经超出了最大值了.感觉后面要不要优化
	fmt.Println(-1^(-1<<7)) //4,0
	fmt.Println(math.MaxInt8)
}
