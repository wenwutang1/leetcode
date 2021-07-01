package _021

import "testing"

func Test32(t *testing.T){
	removeDuplicates([]int{1,1,1,2,2,3})

}

// [1,1,1,2,2,3]
func removeDuplicates(nums []int) int {
	if len(nums)<2{
		return 0
	}
	var i,j int =0,1
	for j< len(nums){
		if nums[j]==nums[i]{
			j++
		}else{
			nums[i+1]=nums[j]
			i++
		}
	}
	return j
}
