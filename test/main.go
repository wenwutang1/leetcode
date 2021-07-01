package main

import (
	"context"
	"fmt"
	"time"
)

var a ="gg"
func main(){


	var wait =make(chan int,0)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	go func(ctx context.Context) {
		tag:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("超时")
				wait<-100
				break tag
			default:
				fmt.Println("芜湖")
				//time.Sleep(time.Second)
			}
		}
	}(ctx)
	//cancel()
	<-wait
	for{
		fmt.Println("main")
		time.Sleep(time.Second*1)
	}
	fmt.Println("GG")
}
