package back

import (
	"fmt"
	"testing"
)

func Test_11(t *testing.T) {
	fmt.Println(permuteUnique([]int{1,2,3}))
}

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列
func permuteUnique(nums []int) [][]int {
	var result = make([][]int, 0)
	var flag []int = make([]int, len(nums))
	backback(nums, &result, []int{}, flag)

	return result
}

func backback(nums []int, result *[][]int, road []int, flag []int) {
	if len(road) == len(nums) {
		var temp = make([]int, len(road))
		copy(temp, road)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if flag[i] == 1 {
			continue
		}
		j:=i-1
		for ;j>-1;j--{
			if nums[i]==nums[j]&&flag[j]==1{
				break
			}
		}
		if j>-1{
			continue
		}
		flag[i] = 1
		backback(nums, result, append(road, nums[i]), flag)
		//同一层相同的节点去掉
		flag[i] = 0
	}
}
