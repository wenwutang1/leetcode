package container

import (
	"fmt"
	"testing"
)

type HeapInt []int

func (h HeapInt) Len() int {
	return len(h)
}

func (h HeapInt) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *HeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Test_heap(t *testing.T){
	excel2String := Excel2String(701)
	fmt.Println(excel2String)
	excel2Int := Excel2Int("ZY")
	fmt.Println(excel2Int)
}


//Excel数字对应的列
// 1->A
// 2->B
// ...
// 28->AB
func Excel2String(col int)(column string){
	if col<1{
		return
	}
	var array []uint8=make([]uint8,0)
	for col>0{
		col--
		array=append(array,uint8('A'+col%26))
		col=col/26
	}
	for i,j:=0, len(array)-1;i<j;i,j=i+1,j-1{
		array[i]=array[i]^array[j]
		array[j]=array[i]^array[j]
		array[i]=array[i]^array[j]
	}
	return string(array)
}
//Excel列对应的数字
// A->1
// B->2
// ...
// AB->28
func Excel2Int(column string)(col int){
	if column==""{
		return 0
	}
	for i:=0;i< len(column);i++{
		col=col*26+int(column[i]-'A'+1)
	}
	return col
}
