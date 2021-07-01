package _021

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_safe1(t *testing.T){

	//unsafe包使用
	//场景1:解决第三方库提供的指针类型所使用的数据类型的不匹配问题
	//场景2:解决数据结构体中，不可见属性的取值和修改操作

	//任何的指针类型都能够被转化为Pointer
	//Pointer可以转化为任何类型的指针
	var vary  int =1000
	var ptr1 *int=&vary
	fmt.Println(ptr1)
	p1 := unsafe.Pointer(&vary)
	fmt.Println(p1)
	//pointer转化为*int指针
	 intPtr:=(*int)(p1)
	 fmt.Println(*intPtr)
	 //转化为float32的指针
	 float32Ptr:=(*int8)(p1)
	 fmt.Println(*float32Ptr)
}


type Hash struct {
	Id byte
	name int32
	age  int8
}

type HashStruct struct {
	name int32
	Id byte
	age  int8
}

func Test_safe2(t *testing.T){
	hash1 := Hash{}
	size1 := unsafe.Sizeof(hash1)
	fmt.Println(size1)

	hash2 := HashStruct{}
	size2 := unsafe.Sizeof(hash2)
	fmt.Println(size2)
}
