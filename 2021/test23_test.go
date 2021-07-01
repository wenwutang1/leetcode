package _021

import (
	"fmt"
	"testing"
)

func Test24(t *testing.T){

	var grid=[][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}

	fmt.Println(maxAreaOfIsland1(grid))

}

//一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
//示例 1:
//[[0,0,1,0,0,0,0,1,0,0,0,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,1,1,0,1,0,0,0,0,0,0,0,0],
// [0,1,0,0,1,1,0,0,1,0,1,0,0],
// [0,1,0,0,1,1,0,0,1,1,1,0,0],
// [0,0,0,0,0,0,0,0,0,0,1,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,0,0,0,0,0,0,1,1,0,0,0,0]]
//对于上面这个给定矩阵应返回6。注意答案不应该是 11

//深度优先遍历，从数组的1开始遍历走过的标记为2 避免重复走入
//由当前的节点向四个方向寻找.
func maxAreaOfIsland(grid [][]int) int {
	if len(grid)==0{
		return 0
	}
	type DFS func([][]int,int,int,*int)
	var dfs DFS
	var max int=-1<<31
	dfs= func(nums [][]int, x int, y int,cur *int)  {
		//边界
		if x<0||x== len(nums)||y<0||y== len(nums[x]){
			return
		}
		if nums[x][y]==0||nums[x][y]==2{
			return
		}
		//统计每个岛屿的面积
		*cur++
		nums[x][y]=2
		dfs(nums,x-1,y,cur)
		dfs(nums,x+1,y,cur)
		dfs(nums,x,y-1,cur)
		dfs(nums,x,y+1,cur)
	}
	for i:=0;i< len(grid);i++{
		for j:=0;j< len(grid[i]);j++{
			if grid[i][j]==1{
				cur:=0
				dfs(grid,i,j,&cur)
				if cur>max{
					max=cur
				}
			}
		}
	}
	return max
}


func maxAreaOfIsland1(grid [][]int) int{
	if len(grid)==0{
		return 0
	}
	type DFS func([][]int,int,int)int
	var dfs DFS
	var max int=-1<<31
	dfs= func(nums [][]int, x int, y int)  int{
		//边界
		if x<0||x== len(nums)||y<0||y== len(nums[x]){
			return 0
		}
		if nums[x][y]==2||nums[x][y]==0{
			return 0
		}
		//统计每个岛屿的面积
		nums[x][y]=2

		return 1+dfs(nums,x-1,y)+dfs(nums,x+1,y)+dfs(nums,x,y-1)+dfs(nums,x,y+1)
	}
	for i:=0;i< len(grid);i++{
		for j:=0;j< len(grid[i]);j++{
			if grid[i][j]==1{
				cur:=dfs(grid,i,j)
				if cur>max{
					max=cur
				}
			}
		}
	}
	return max

}
