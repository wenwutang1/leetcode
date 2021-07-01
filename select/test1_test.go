package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)


//限制最的运行次数,最多可以同时运行
var work =make(chan uint8,2)

func TestStr(t *testing.T) {
	
	var call =map[string]func(key string){
		"1": func(str string) {
			fmt.Println(str)
		},
		"2": func(key string) {
			fmt.Println(key)
		},
		"3": func(key string) {
			fmt.Println(key)
		},
	}

	for val,fn:=range call{

		go func() {
			work<-1
			fn(val)
			<-work
		}()
	}
	for {

	}
}


//字符串强制转换为[]slice
//底层的数据接口[]slice多出来一个len的字段
//
func String2Bytes(str string)[]byte{
	//拿到字符串原来的指针类型

	sh:=(*reflect.StringHeader)(unsafe.Pointer(&str))

	bh:=reflect.SliceHeader{
		//返回的slice没有重新的Memory
		Data: sh.Data,
		Len: sh.Len,
		Cap: sh.Len,
	}
	return *(*[]byte) (unsafe.Pointer(&bh))
}

