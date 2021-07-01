package tw

import (
	"fmt"
	"testing"
)

//统计一个数字在排序数组中出现的次数
//nums = [5,7,7,8,8,10], target = 8 输出 2
//[5,7,7,8,8,10], target = 6        输出 0

func Test_ts(t *testing.T){
	search([]int{2,2},2)
}

//找到比目标值小的最大的一个数
func search(nums []int, target int) int {

	l,r,sum:=0, len(nums)-1,0
	for l<=r{
		mid:=int(uint(l+r)>>1)
		if nums[mid]>target{
			r=mid-1
		}else if nums[mid]<target{
			l=mid+1
		}else{
			sum++
			var in1,in2=mid-1,mid+1
			for (in1>-1&&nums[in1]==target)||(in2< len(nums)&&nums[in2]==target){
				if in1>-1&&nums[in1]==target{
					sum++
				}
				if in2< len(nums)&&nums[in2]==target{
					sum++
				}
				in1--
				in2++
			}
			break
		}
	}
	fmt.Println(sum)
	return sum
}
