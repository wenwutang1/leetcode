package tw

import (
	"fmt"
	"testing"
)

//把一个数组最开始的若干个元素搬到数组的末尾,我们称之为数组的旋转
//输入一个递增排序的数组的一个旋转,输出旋转数组的最小元素
//例如,数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1
func Test_t2(t *testing.T) {
	fmt.Println(minArray([]int{3,1,3,3}))
}
func minArray(numbers []int) int {
	return minArray2(numbers, 0, len(numbers)-1)
}

func minArray2(numbers []int, l, r int) int {
	if l+1 >= r {
		return min(numbers[l], numbers[r])
	}
	var mid int = l + (r-l)>>1
	if numbers[l]<numbers[r]{
		return numbers[l]
	}
	//分别二分左右两边
	return min(minArray2(numbers,l,mid-1),minArray2(numbers,mid,r))
}

func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}





