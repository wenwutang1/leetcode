package _021

import (
	"fmt"
	"sort"
	"testing"
)

type student struct {
	age    int
	name   string
	height float64
}
type stuSort func(p1, p2 *student) bool
type studentSort struct {
	data []student //需要排序的数据
	less []stuSort //排序的方法
}

func (this *studentSort) Len() int {
	return len(this.data)
}

func (this *studentSort) Swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this *studentSort) Less(i, j int) bool {
	p, q := &this.data[i], &this.data[j]
	var k int
	for k = 0; i < len(this.less)-1; k++ {
		less := this.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
		//default
		//相等会进入下一个策略
	}
	return this.less[k](p, q)
}

func (this *studentSort) Sort(data []student) {
	this.data = data
	sort.Sort(this)
}

func OrderBy(less ...stuSort) *studentSort {
	return &studentSort{
		less: less,
	}
}

func Test_test4(t *testing.T) {
	var change = []student{
		{name: "沸羊羊", age: 18, height: 15.5},
		{name: "喜洋洋", age: 27, height: 18.5},
		{name: "美羊羊", age: 28, height: 17.5},
		{name: "黑羊羊", age: 28, height: 16.5},
		{name: "大山羊", age: 27, height: 19.00},
	}
	name := func(p1, p2 *student) bool {
		return p1.name < p2.name
	}
	age := func(p1, p2 *student) bool {
		return p1.age < p2.age
	}
	height := func(p1, p2 *student) bool {
		return p1.height < p2.height
	}
	//只根据名字排序
	OrderBy(name).Sort(change)
	fmt.Println("根据名字排序:", change)
	//先根据年龄排序,再根据身高排序
	OrderBy(age, height).Sort(change)
	fmt.Println("根据年龄排序,相同的根据身高排序:", change)
	//先根据名字,年龄排序,再根据身高排序
	OrderBy(name, age, height).Sort(change)
}

//BUG(tww):memory leak
func Test_t5(t *testing.T){

	var slice []int
	fmt.Println(slice==nil)
	slice=*new([]int)
	fmt.Println(slice==nil)

}