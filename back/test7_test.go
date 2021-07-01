package back

import (
	"fmt"
	"testing"
)

func Test_02(t *testing.T) {

	//var grid = [][]int{
	//	{1,6,1},
	//	{5,8,7},
	//	{1,9,1},
	//}
	//gold := getMaximumGold(grid)
	//fmt.Println(gold)
	//numTilePossibilities("AAB")
	//ints := intersection([]int{1, 2, 2, 1}, []int{2,2})
	//fmt.Println(ints)
	//generateSort(5678)
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}

//你要开发一座金矿，地质勘测学家已经探明了这座金矿中的资源分布，并用大小为
//m * n 的网格 grid 进行了标注。每个单元格中的整数就表示这一单元格中的黄金数量
//如果该单元格是空的，那么就是 0
//为了使收益最大化，矿工需要按以下规则来开采黄金：
//每当矿工进入一个单元，就会收集该单元格中的所有黄金。
//矿工每次可以从当前位置向上下左右四个方向走。
//每个单元格只能被开采（进入）一次。
//不得开采（进入）黄金数目为 0 的单元格。
//矿工可以从网格中 任意一个 有黄金的单元格出发或者是停止
//grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
//[[1,0,7],
// [2,0,6],
// [3,4,5],
// [0,3,0],
// [9,0,20]]
//一种收集最多黄金的路线是：1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7。
func getMaximumGold(grid [][]int) int {
	//1 收集的最大的数量
	var max = 0
	var flag = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		flag[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				//作为入口
				backji(grid, flag, 0, i, j, &max)
			}
		}
	}
	return max
}

func backji(grid [][]int, flag [][]int, cur int, l int, r int, result *int) {
	//路径为0的不能通过
	if l >= len(grid) {
		return
	}
	if r >= len(grid[l]) {
		return
	}
	if flag[l][r] == 1 {
		if *result < cur {
			*result = cur
		}
		return
	}
	if grid[l][r] == 0 {
		//路径已经走完了比较大小
		if *result < cur {
			*result = cur
		}
		return
	}
	flag[l][r] = 1
	//四个方向行走
	if l < len(grid)-1 {
		//向下走
		backji(grid, flag, cur+grid[l][r], l+1, r, result)
	}
	//向上走
	if l > 0 {
		backji(grid, flag, cur+grid[l][r], l-1, r, result)
	}
	//向左走
	if r > 0 {
		backji(grid, flag, cur+grid[l][r], l, r-1, result)
	}
	if r < len(grid[l])-1 {
		backji(grid, flag, cur+grid[l][r], l, r+1, result)
	}
	//四个方向都走完了,回退到上一步
	flag[l][r] = 0
}

//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//定义flag
//输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"
func numTilePossibilities(tiles string) int {
	var flag = make([]bool, len(tiles))
	var result = 0
	backhz(tiles, flag, "", &result)
	return result
}

func backhz(tiles string, flag []bool, road string, sum *int) {
	for i := 0; i < len(tiles); i++ {
		if flag[i] {
			continue
		}
		//跳过同层级的节点
		var j = i - 1
		for ; j > -1; j-- {
			if tiles[i] == tiles[j] && flag[j] == false {
				break
			}
		}
		if j > -1 {
			continue
		}
		flag[i] = true
		*sum += 1
		backhz(tiles, flag, road+string(tiles[i]), sum)
		flag[i] = false
	}
}

func intersection(nums1 []int, nums2 []int) []int {

	mmp := make(map[int]int8)
	var result = make([]int, 0)
	//小数组驱动大数组
	var i = 0
	for i = 0; i < len(nums1); i++ {
		mmp[nums1[i]] = 1
	}
	for i = 0; i < len(nums2); i++ {
		if v, ok := mmp[nums2[i]]; ok && v == 1 {
			result = append(result, nums2[i])
			mmp[nums2[i]]++
		}
	}
	return result
}

//我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。
//请你返回由[low, high]范围内所有顺次数组成的 有序 列表（从小到大排序）。
//输出：low = 100, high = 300
//输出：[123,234]
//输出：low = 1000, high = 13000
//输出：[1234,2345,3456,4567,5678,6789,12345]
func sequentialDigits(low int, high int) []int {
	return nil
}
func generateSort(target int) int {
	var result = target % 10
	var t = result
	if t == 9 {
		t = 1
	}
	for target > 10 {
		target = target / 10
		t++
		result = result*10 + (t)

	}
	fmt.Println(result)
	return 0
}

//给定一个字符串数组 arr，字符串 s 是将 arr 某一子序列字符串连接所得的字符串，如果 s 中的每一个字符都只出现过一次，那么它就是一个可行解。
//请返回所有可行解 s 中最长长度。
//输入：arr = ["un","iq","ue"]
//输出：4
//解释：所有可能的串联组合是 "","un","iq","ue","uniq" 和 "ique"，最大长度为 4。
//输入：arr = ["cha","r","act","ers"]
//输出：6
//解释：可能的解答有 "chaers" 和 "acters"。
func maxLength(arr []string) int {
	var max = 0
	Back123(arr, 0, "", &max)
	return max
}

func Back123(nums []string, index int, road string, max *int) {
	if index == len(nums) {
		return
	}
	if !repeat(road) {
		return
	}
	for i := index; i < len(nums); i++ {
		//本身就是重复的字符串,直接跳过
		if !repeat(road + nums[i]) {
			continue
		}
		road = road + nums[i]
		if *max < len(road) {
			*max = len(road)
		}
		Back123(nums, i+1, road, max)
		if *max == 26 {
			return
		}
		road = road[:len(road)-len(nums[i])]
	}
}

//检查一个字符串是否有重复的字符
func repeat(astr string) bool {
	var mark int = 0
	for i := 0; i < len(astr); i++ {
		offset := astr[i] - 'a'
		tmp := 1 << offset
		if (mark & tmp) == tmp {
			return false
		} else {
			mark |= tmp
		}
	}
	return true
}
