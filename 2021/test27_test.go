package _021

import "testing"

func Test28(t *testing.T){

}

//30. 被围绕的区域
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
//
//示例 1： x x x x       x x x x
//        x o o x       x x x x
//        x o x x  ====>x x x x
//        x x o x       x x o x
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//示例 2：



//结题思路，深度优先遍历.由于要找到所有被x包围的哦，并且把o变为x.
//并且这个o的集合不能与边界的o相连.
//第一步 找到所有与边界的o相连的o,将他们设置为'A'.
//第二部遍历整个集合,把所有的o设置为x
//第三部把设置了为o的'A'还原为'O'
func solve(board [][]byte)  {
	if len(board)<1{
		return
	}
	type dtf1 func([][]byte,byte,int,int)
	var  border dtf1
	border= func(nums [][]byte, flag byte, x int, y int) {
		if x<0||x>= len(nums)||y<0||y>= len(nums[x])||nums[x][y]=='x'{
			return
		}
		if flag=='o'{
			if nums[x][y]=='A'{
				return
			}
		}
		if flag=='A'{
			if nums[x][y]=='0'{
				return
			}
		}
		if flag=='x'{
			nums[x][y]='x'
		}
		nums[x][y]=flag
		border(nums,flag,x,y-1)
		border(nums,flag,x,y+1)
		border(nums,flag,x-1,y)
		border(nums,flag,x+1,y)
	}
	//第一部函数
	for j:=0;j< len(board[0]);j++{
		border(board,'A',0,j)
		border(board,'A',1,j)
	}
	//第二部把所有的内部包围o换成x
	for i:=0;i< len(board);i++{
		for j:=0;j< len(board[i]);j++{
			if board[i][j]=='o'{
				border(board,'x',i,j)
			}
		}
	}
	//最后把A换回来
	for j:=0;j< len(board[0]);j++{
		border(board,'o',0,j)
		border(board,'o',1,j)
	}
}

