package tw

import (
	"fmt"
	"testing"
)

//给你一个排序后的字符列表 letters ,列表中只包含小写英文字母。
//另给出一个目标字母target，请你寻找在这一有序列表里比目标字母大的最小字母。
//在比较时，字母是依序循环出现的。举个例子
//如果目标字母 target = 'z' 并且字符列表为letters = ['a', 'b']，则答案返回'a'
func Test_t3(t *testing.T) {
	fmt.Println(string(nextGreatestLetter3([]byte{'c', 'f', 'j'}, 'k')))
}
func nextGreatestLetter(letters []byte, target byte) byte {
	return nextGreatestLetter2(letters, target, 0, len(letters)-1)
}
//二分法
func nextGreatestLetter2(letters []byte, target byte, l, r int) byte {
	if l > r {
		return letters[l%len(letters)]
	}
	var mid int = int(uint(l+r) >> 1)
	//小于这个数不用进度左边的分治
	if letters[mid] > target {
		return nextGreatestLetter2(letters, target, l, mid-1)
	} else {
		return nextGreatestLetter2(letters, target, mid+1, r)
	}
}


func nextGreatestLetter3(letters []byte, target byte) byte {
	l,r:=0, len(letters)-1
	for l<=r{
		mid:=l+(r-l)>>1
		if letters[mid]>=target{
			l=mid+1
		}else if letters[mid]<target{
			r=mid-1
		}
	}
	return letters[l%len(letters)]
}