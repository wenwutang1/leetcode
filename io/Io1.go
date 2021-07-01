package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

//////////////////////////同步阻塞IO
func blockRead(){
	//read 读取文件从用户态到内核态->处理一直等待的状态知道操作系统文件读取完毕通知线程
	//线程才能恢复只执行
	file,err:=os.OpenFile("./io/g1.txt",syscall.O_RDONLY,1)
	if err!=nil{
		return
	}
	defer file.Close()
	bytes:=make([]byte,1024)
	buff:=make([]byte,0)
	for {
		n,err:=file.Read(bytes)
		if err!=nil&&err!=io.EOF{
			panic("fatal error")
		}else if err==io.EOF{
			break
		}
		buff=append(buff,bytes[0:n]...)
	}
	fmt.Println(string(buff))
}
//////////////////////////同步非阻塞IO
//实现方式采用状态轮询->状态通知方式
//
func noBlockRead(){
	//定时器
	//每间隔两秒查看文件是否准备就绪
	var state chan struct{} =make(chan struct{},0)
	go func() {
		//模拟读写文件的操作
		time.Sleep(time.Second*2)
		state<- struct{}{}
		close(state)
	}()
	var loop bool
	for {
		select {
		case _,ok:=<-state:
			if !ok{
				fmt.Println("文件读取完毕")
				loop=true
			}
		}
		if loop{
			break
		}
	}
}
func main(){
	ch := make(chan string)
	go func(ch chan string) {
		//模拟读取文件的操作
		time.Sleep(time.Second*10)
		ch <- "A line of text"
		close(ch)
	}(ch)

stdinloop:
	for {
		select {
		case stdin, ok := <-ch:
			if ok {
				fmt.Println("Read input from stdin:", stdin)
				break stdinloop

			}
		case <-time.After(1 * time.Second):
			fmt.Println(" Do something when there is nothing read from stdin")
		}
	}
	fmt.Println("Done, stdin must be closed")
	var i int8 =0
	loop:
	for {
		i++
		fmt.Println(i)
		if i==10{
			break loop
		}
	}
	fmt.Println("ad")
}



