package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main(){
	wait:=sync.WaitGroup{}

	arr:=make([]demo,0,100)
	for i:=0;i<20;i++{
		arr=append(arr,demo{
			id: i+1,
			name: strconv.Itoa(i)+"name",
		})
	}
	wait.Add(len(arr))
	for _,order:=range arr{
		go func(r demo) {
			fmt.Println(r.name)
			wait.Done()
		}(order)
		//go func() {
		//	fmt.Println(order.name)
		//	wait.Done()
		//}()
	}
	wait.Wait()
	fmt.Println("-------------------------------")
}


type demo struct {
	id int
	name string
}
