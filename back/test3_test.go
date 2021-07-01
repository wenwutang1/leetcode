package back

import (
	"sort"
	"testing"
)

func Test_O(t *testing.T){
	kClosest([][]int{{1,3},{-2,2},{2,-2}},2)
}
//我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
//（这里，平面上两点之间的距离是欧几里德距离。）
func kClosest(points [][]int, K int) [][]int {

	node := make([]struct {
		x   int
		y   int
		dis int
	}, 0, len(points))

	for _, v := range points {
		node = append(node, struct {
			x   int
			y   int
			dis int
		}{x: v[0], y: v[1], dis: v[0]*v[0] + v[1]*v[1]})
	}
	sort.Slice(node, func(i, j int) bool {
		return node[i].dis < node[j].dis
	})
	ans := make([][]int, 0)
	for i := 0; i < K; i++ {
		ans = append(ans, []int{node[i].x,node[i].y})
	}
	return ans
}
