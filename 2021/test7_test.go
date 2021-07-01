package _021

import (
	"fmt"
	"testing"
)

func Test77(t *testing.T){
	fmt.Println(restoreIpAddresses("25525511135"))
}

//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
func restoreIpAddresses(s string) []string {

	result := make([]string, 0,800)
	restoreIpAddresses2(s,0,"",[]string{},result)
	return result
}

func restoreIpAddresses2(s string,index int,road string,temp []string,result []string){

}


