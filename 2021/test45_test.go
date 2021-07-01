package _021

import (
	"fmt"
	"sort"
	"testing"
)

func Test46(t *testing.T) {
	fmt.Println(intersection2([]int{4,5,9}, []int{4,4,8,9,9}))
}

func intersection(nums1 []int, nums2 []int) []int {
	var mmp = make(map[int]byte)
	for i := range nums1 {
		mmp[nums1[i]] = 1
	}

	var res = make([]int, 0)
	for i := range nums2 {
		if _, ok := mmp[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(mmp, nums2[i])
		}
	}
	return res
}

//nums1 = [4,9,5], nums2 = [9,4,9,8,4]

//[4,5,9]  [4,4,8,9,9]
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	var result = make([]int, 0)
	var l1 = len(nums1)
	var l2 = len(nums2)
	var index1, index2 = 0, 0
	for index1 < l1 && index2 <l2  {
		if nums1[index1] == nums2[index2] {
			if len(result) == 0 || (len(result) > 0 && nums1[index1] != result[len(result)-1]) {
				result = append(result, nums1[index1])
			}
			index1++
			index2++
			continue
		}
		if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}

	return result
}
