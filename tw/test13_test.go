package tw

import (
	"fmt"
	"testing"
)

//输入：
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出：3
//解释：
//长度最长的公共子数组是 [3, 2, 1]
func Test_find(t *testing.T) {

	//length := findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
	//fmt.Println(length)
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

//动态规划
//			1 2 3 2 1
//     3    0 0 1 0 0
//     2    0 1 0 1 0
//     1    1 0 0 0 1
//     4    0 0 0 0 0
//     7    0 0 0 0 0
//如上图解决动态规划的问题在于构件动态数组找到状态转移的方程
//联系子数组的最大长度应该是所有的1唯一一条斜线上
//dp[i][j]表示当前 A[i]B[j]对应的联系数组的长度
//dp[i][j]= {if A[i]=A[j]=1+dp[i-1][j-1]}
func findLength(A []int, B []int) int {
	dp := make([][]int, len(A))
	var max = 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(B))
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] && i > 0 && j > 0 {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else if A[i] == B[j] {
				dp[i][j] = 1
			}
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	//返回dp[i][j]中的最大的数组
	return max
}

//如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。
//例如，[1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3)是正负交替出现的。相反, [1,4,7,2,5]和[1,7,4,5,5] 不是摆动序列，
//第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
//给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
//输入: [1,7,4,9,2,5]
//输出: 6
//解释: 整个序列均为摆动序列。
//示例 2:
//输入: [1,17,5,10,13,15,10,5,16,8]
//输出: 7
//dp[i][0]上升
//dp[i][1]下降
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] > nums[i-1] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][0] + 1
		} else if i > 0 && nums[i] < nums[i-1] {
			dp[i][0] = dp[i-1][1] + 1
			dp[i][1] = dp[i-1][1]
		} else if i > 0 {
			//相当于删除,不能连续的数字
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1]
		}
	}
	return maxI(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func maxI(a, b int) int {
	if a < b {
		return b
	}
	return a
}
