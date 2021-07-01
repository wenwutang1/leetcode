package _021

import "testing"

func Test18(t *testing.T){


}

//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数

//第一数值是排序好了的
//1 对每一行使用二分查找法
//2 找到了就直接返回
//没找到和每一行的开头结尾比较
func searchMatrix(matrix [][]int, target int) bool {
	var index int
	for index>0&&index< len(matrix){
		var array=matrix[index]
		//二分查找法
		var l,r=0, len(array)-1
		for l<r{
			mid:=(l+r)>>1
			if array[mid]==target{
				return true
			}
			if array[mid]<target{
				r=mid-1
			}else{
				l=mid+1
			}
		}
		if target<array[0]{
			index--
		}else if target>array[len(array)-1]{
			index++
		}
	}
	return false
}