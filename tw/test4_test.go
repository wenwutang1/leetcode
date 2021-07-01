package tw

import (
	"testing"
)

func Test_03(t *testing.T){
//[[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
	countNegatives([][]int{
		{3,-1,-3,-3,-3},
		{2,-2,-3,-3,-3},
		{1,-2,-3,-3,-3},
		{0,-3,-3,-3,-3},
	})

}

//给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列
func countNegatives(grid [][]int) int {
	sum:=0
	for i:=0;i< len(grid);i++{
		//找到比0,小的最大的值这个节点,最后R
		l,r:=0, len(grid[i])-1
		if grid[i][r]>=0{
			continue
		}
		if grid[i][0]<0{
			sum+=r+1
			continue
		}
		for l<=r{
			mid:=int(uint(l+r)>>1)
			if grid[i][mid]<0{
				r=mid-1
			}else{
				l=mid+1
			}
		}
		sum+= len(grid[i])-l
	}
	return sum
}