package main

import (
	"fmt"
	"sync"
	"testing"
)


var globalEnv string =""
var once sync.Once
func OnlyDoOne(){
	globalEnv="我在干啥么"
	fmt.Println(globalEnv)
}
func TestOnce(t *testing.T){

	var wait =make(chan uint8,2)

	go func(){
		once.Do(OnlyDoOne)
		wait<-10
	}()
	go func(){
		once.Do(OnlyDoOne)
		wait<-11
	}()
	<-wait
	<-wait

}
