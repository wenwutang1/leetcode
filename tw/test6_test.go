package tw

import (
	"fmt"
	"strings"
	"testing"
)

var target int64=0


func Test_GO(t *testing.T){

	//var wait sync.WaitGroup
	//wait.Add(100000)
	//for i:=0;i<100000;i++{
	//	go func() {
	//		atomic.AddInt64(&target,1)
	//		wait.Done()
	//	}()
	//}
	//wait.Wait()
	//fmt.Println(target)

	str:="777.77"
	fmt.Println(strings.TrimSpace(str))
}
