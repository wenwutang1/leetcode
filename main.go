package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

//go chan 原理
//chan 使用
//使用make创建,可以指定长度.如果不指定那么默认就是0
//0代表缓冲区buf大小为0,发送和接受都会阻塞当前的线程,必须要有其他的线程参与发送和接受
//>0代表有缓冲区,发送的数据数量<=缓冲区的大小,不会阻塞当前的线程.

//chan在创建时可以指定为可发送chan,可接受chan,既可以发送又可以接受(默认make(chan T))

//chan的关闭方法内置函数close方法,向一个close的chan发送消息会引起panic
//chan的遍历@1
func main(){


	//ch:=make(chan struct{},0)
	////不会阻塞
	//
	////创建一个只能发送的chan
	////ch1:=make(chan<- struct{},0)
	////ch1<- struct{}{}
	//
	////创建一个只能接受的chan
	////ch2:=make(<-chan struct{},0)
	////编译报错因为这个chan只能接受
	////ch2<- struct{}{}
	//
	////@1chan的遍历
	//
	//go func() {
	//	ch<- struct{}{}
	//	fmt.Println("放入一个东西")
	//	//不关闭chan
	//	close(ch)
	//}()
	//
	////使用range 的方式
	////range的方式也会阻塞(除非这个chan被close了)
	////对一个被close的chan,接受数据会不会阻塞
	//for k:=range ch{
	//	fmt.Println(k)
	//}
	////对已经	关闭的chan读操作,不会阻塞可以循环读取.
	//a:=<-ch
	//fmt.Println(a)
	Test_3g(100)

}


func string2Byte(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}


//chan的基本操作,使用3个groutine
func Test_3g(count int){
	wait:=sync.WaitGroup{}
	wait.Add(3)
	ch1:=make(chan struct{},0)
	ch2:=make(chan struct{},0)
	ch3:=make(chan struct{},0)
	go func() {
		defer wait.Done()
		c:=0
		for {
			select {
			case <-ch1:
				fmt.Println("fish")
				c++
				ch2<- struct{}{}
			}
			if c==count{
				close(ch1)
				return
			}
		}
	}()

	go func() {
		c:=0
		defer wait.Done()
		for {
			select {
			case <-ch2:
				fmt.Println("cat")
				ch3<- struct{}{}
				c++
			}
			if c==count{
				close(ch2)
				return
			}
		}
	}()

	go func() {
		defer wait.Done()
		c:=0
		for {
			select {
			case <-ch3:
				fmt.Println("dog")
				c++
				if c==count{
					close(ch3)
					return
				}
				ch1<- struct{}{}
			}
		}
	}()
	//首先打印fish
	ch1<- struct{}{}
	wait.Wait()
	fmt.Println("main end")
}


//预备知识
//goroutine 知识
// type g struct{
//	atomicstatus uint32            //表示goroutine的状态
//	param        unsafe.Pointer   //唤醒时参数
//  waitreason   waitReason       //等待原因
//	waiting        *sudog         //表示当前的wait正在等待谁活着正在等待的事务.
//}
//
//func getg() *g 获取当前的g对象.

//sudog对象
//type sudog struct {
// 	isSelect bool
//  elem     unsafe.Pointer
//  waitlink    *sudog // 连接的下一个对象
// 	c           *hchan //
//
//}