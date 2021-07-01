package _021

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T){

	var grid =[][]byte{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
	fmt.Println(numIslands(grid))
}

//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
//示例 1：
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1


//解决思路，深度优先遍历。往四个方向遍历，把走过的的1变为0，深度优先遍历的次数即为岛屿的数量
func numIslands(grid [][]byte) int {
	if len(grid)<1{
		return 0
	}
	type DFS func([][]byte,int,int)
	var dfs DFS
	dfs= func(nums [][]byte, x int, y int) {
		//边界返回
		if x<0|| x== len(nums)||y<0||y== len(nums[x]){
			return
		}
		//遇到水无法连接岛屿
		//标记走过了的陆地不能往返递归
		if nums[x][y]=='0' ||nums[x][y]=='2'{
			return
		}
		nums[x][y]='2'
		dfs(nums,x-1,y)
		dfs(nums,x+1,y)
		dfs(nums,x,y-1)
		dfs(nums,x,y+1)
	}
	var (
		result int
	)
	for i:=0;i< len(grid);i++{
		for j:=0;j< len(grid[i]);j++{
			if grid[i][j]=='1'{
				result++
				dfs(grid,i,j)
			}
		}
	}
	return result
}
