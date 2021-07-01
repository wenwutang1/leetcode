package tw

import "testing"

func Test_is(t *testing.T){

}

func isBadVersion(version int) bool{
	return false
}

func firstBadVersion(n int) int {

	l,r:=0,n-1
	for l<=r{
		mid:=int(uint(l+r)>>1)
		if isBadVersion(mid){
			r=mid-1
		}else{
			l=mid+1
		}
	}
	return l
}


