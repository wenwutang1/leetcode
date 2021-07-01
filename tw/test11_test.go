package tw

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func Test_t111(t *testing.T) {
	//queue := reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
	//fmt.Println(queue)
	fmt.Println(removeKdigits("1234567890", 9))

}

//假设有打乱顺序的一群人站成一个队列 每个人由一个整数对(h, k)表示其中h是这个人的身高
//k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
//根据身高排序
//再根据k排序
//[7,0],[7,1],[6,1],[5,0][5,2][4,4]
//第一步
//[7,0]
//[7,0],[7,1]
//[7,0],[6,1],[7,1]
//[5,0],[7,0],[6,1],[7,1]
//[5,0],[7,0],[5,2],[6,1],[7,1]
//[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]
func reconstructQueue(people [][]int) [][]int {
	var result = make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		result[i] = make([]int, 2)
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})
	var j = 0
	for i := 0; i < len(people); i++ {
		for j = len(result) - 1; j > people[i][1]; j-- {
			result[j] = result[j-1]
		}
		result[people[i][1]] = people[i]
	}
	return result
}

//["flower","flow","flight"]
//寻找最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	//找到最大的子串和最小的子串
	var max, min string = strs[0], strs[0]
	for i := 0; i < len(strs); i++ {
		if strs[i] < min {
			min = strs[i]
		}
		if strs[i] > max {
			max = strs[i]
		}
	}
	var result []byte = make([]byte, 0, len(max))
	var i = 0
	for i < len(min) && i < len(max) && min[i] == max[i] {
		result = append(result, max[i])
		i++
	}
	return string(result)
}

//给定一个以字符串表示的非负整数 num,移除这个数中的 k 位数字,使得剩下的数字最小
//输入: num = "1432219", k = 3
//输出: "1219"
//解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
//输入: num = "10200", k = 1
//输出: "200"
//解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//num = "10", k = 2
//输出: "0"
//暴力破解-超出时间限制
func removeKdigits(num string, k int) string {
	var flag = make([]byte, len(num))
	var result = -1
	removeKdigits2(num, flag, 0, k, &result)
	return strconv.Itoa(result)
}
func removeKdigits2(num string, flag []byte, c int, k int, result *int) {
	if c == k {
		road := make([]byte, 0, len(num)-k)
		for i := 0; i < len(num); i++ {
			if flag[i] == 0 {
				road = append(road, num[i])
			}
		}
		str := string(road)
		parseInt, _ := strconv.ParseInt(str, 10, 64)
		if *result == -1 || *result > int(parseInt) {
			*result = int(parseInt)

		}
		return
	}
	for i := 0; i < len(num); i++ {
		if flag[i] == 1 {
			continue
		}
		//移除
		flag[i] = 1
		//如果首位明显小于前一个结果不用跳过
		if *result!=-1&&c==0&&num[c]>strconv.Itoa(*result)[0]{
			continue
		}
		removeKdigits2(num, flag, c+1, k, result)
		flag[i] = 0
	}
}

//贪心算法
func removeKdigits3(num string, k int) string {

	//var result =make([]byte,0, len(num)-k)
	//var flag   =make([]uint8,0, len(num))
	//var n= 0
	//for n<k{
	//	for j:=0;j< len(flag);j++{
	//		//已经被删除了
	//		if flag[j]==1{
	//			continue
	//		}
	//		if flag[j]<flag[j+1]
	//	}
	//}
	return ""
}