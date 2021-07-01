package tw

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_dis(t *testing.T) {
	fmt.Println(string(1+'0'))
	fmt.Println(summaryRanges([]int{-1}))
}

//[0,1,2,4,5,7]
func summaryRanges(nums []int) []string {
	if len(nums)==0{
		return []string{}
	}
	var result=make([]string,0)
	var ptr1=0
	for ptr1< len(nums){
		var index=ptr1
		for index< len(nums)-1&&nums[index]+1==nums[index+1]{
			index++
		}
		if index!=ptr1{
			result=append(result,strconv.Itoa(nums[ptr1])+"->"+strconv.Itoa(nums[index]))
		}else{
			result=append(result,strconv.Itoa(nums[index]))
		}

		ptr1=index+1
	}
	return result
}

//func convertToBase7(num int) string {
//	if num == 0 {
//		return "0"
//	}
//	var c = num
//	var result string
//	for num != 0 {
//		if num > 0 {
//			result = string((num%7)+'0') + result
//		} else {
//			result = string((-num%7)+'0') + result
//		}
//		num = num / 7
//	}
//	if c < 0 {
//		return "-" + result
//	}
//	return result
//}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}
	if start > destination {
		start, destination = destination, start
	}
	var l, r = 0, 0
	var c = start
	for ; c < destination; c++ {
		l += distance[c]
	}
	for ; destination%len(distance) != start; destination++ {
		r += distance[destination%len(distance)]
	}
	return min(l, r)

}
