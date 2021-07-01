package _021

import (
	"fmt"
	"testing"
)

func Test27( t*testing.T){
	fmt.Println(findOrder(3,[][]int{{1,0}}))
}

//有向图无环图的应用
//又向图的两个基本概念 入度:由其他节点指向当前节点的数量
//       ① ->② 如图 节点1 指向节点2 节点1没有其他的节点指向节点1的入度为0 出度为1
//       节点2 的入度为1 出度为0
//
//拓扑排序是把有向图按照入度的顺序关系进行排序的过程
//拓扑排序主要解决一种顺序关系的模型列，如课程学习[在学习某些高级课程之前必须要学习某些基础的课程]

//1 课程的安排
//207 在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程  bi 。
//
//例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//示例 1：

//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//示例 2：
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
func canFinish(numCourses int, prerequisites [][]int) bool {
	//队列用来存放当前能够学习的课程下标
	var queue []int=make([]int,0)
	//存放有向图的关系
	var mmp =make(map[int][]int)
	//初始化课程下标,存放当前课程的入度
	var course =make([]int,numCourses)
	//遍历课程关系图，设置有向关系图和入度
	for i:=0;i< len(prerequisites);i++{
		//先要学习的课程
		pre:=prerequisites[i][1]
		//后学习的课程
		next:=prerequisites[i][0]
		course[prerequisites[i][0]]++
		if mmp[pre]==nil{
			mmp[pre]=make([]int,0)
		}
		//建立依赖关系，列如[1,0],方便学习完pre课程后对next课程减少入度
		mmp[pre]=append(mmp[pre],next)
	}

	//找到入度为0的课程加入到队列中
	for j:=0;j< len(course);j++{
		if course[j]==0{
			queue=append(queue,j)
		}
	}
	var count = 0
	for len(queue)>0{
		count++
		node:=queue[len(queue)-1]
		queue=queue[:len(queue)-1]
		conn:=mmp[node]
		if conn!=nil{
			//把依赖它的课程号的入度-1
			for index:=0;index< len(conn);index++{
				course[conn[index]]--
				if course[conn[index]]==0{
					//加入到队列
					queue=append(queue,conn[index])
				}
			}
		}
	}
	return count==numCourses
}


//找出一种可能的学习课程的顺序
func findOrder(numCourses int, prerequisites [][]int) []int {

	var queue=make([]int,0)
	var course =make([]int,numCourses)
	var connMap =make(map[int][]int)

	for i:=0;i< len(prerequisites);i++{
		pre:=prerequisites[i][1]
		next:=prerequisites[i][0]
		course[next]++
		if connMap[pre]==nil{
			connMap[pre]=make([]int,0)
		}
		connMap[pre]=append(connMap[pre],next)
	}
	for i:=0;i< len(course);i++{
		if course[i]==0{
			queue=append(queue,i)
		}
	}
	var result =make([]int,0)
	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]
		result=append(result,node)
		list:=connMap[node]
		if list!=nil{
			for index:=0;index< len(list);index++{
				course[list[index]]--
				if course[list[index]]==0{
					queue=append(queue,list[index])
				}
			}
		}
	}

	if len(result)==numCourses{
		return result
	}

	return []int{}
}