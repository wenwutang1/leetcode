package _021

import (
	"fmt"
	"testing"
)

func Test23(t *testing.T) {

	var grid=[][]int{
		{1,0},

	}
	fmt.Println(islandPerimeter(grid))

}

//网格中的格子 水平和垂直 方向相连（对角线方向不相连）
//整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）
//岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）
//格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长
//输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var length int
	type DFS func([][]int, int, int)
	var dfs DFS
	dfs = func(nums [][]int, x int, y int) {
		//是否越界
		if x >= len(nums) || x < 0 || y < 0 || y >= len(nums[x]) ||nums[x][y]==0{
			length++
			return
		}
		if nums[x][y]==2{
			return
		}
		//走过的数组设置为2避免重复
		nums[x][y] = 2
		dfs(nums, x-1, y)
		dfs(nums, x+1, y)
		dfs(nums, x, y-1)
		dfs(nums, x, y+1)
	}
	for i:=0;i< len(grid);i++{
		for j:=0;j< len(grid[i]);j++{
			if grid[i][j]==1{
				dfs(grid,i,j)
				break
			}
		}
	}
	return length
}
