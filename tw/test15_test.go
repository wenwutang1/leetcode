package tw

import (
	"fmt"
	"sort"
	"testing"
)

// [1,3,5,6], 5
//[1,3,5,6], 2
func Test_tt(t *testing.T){
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func countAndSay(n int) string {
	var str ="1"
	var c =1
	for ;c<n;c++{
		start,end:=0,0
		var temp =make([]uint8,0, len(str))
		for start< len(str){
			if str[start]==str[end]{
				end++
			}else if end< len(str){
				//不等于
				temp=append(temp,uint8(end-start)+'0')
				temp=append(temp,str[start])
				//temp+=string((end-start)+'0')+string(str[start])
				start=end
			}
			if end==len(str){
				temp=append(temp,uint8(end-start)+'0')
				temp=append(temp,str[start])
				break
			}
		}
		str=string(temp)
	}
	return str
}

func addBinary(a string, b string) string {
	result := ""
	flag:=0
	t1,t2:= len(a)-1, len(b)-1

	for t1>=0||t2>=0{
		r1,r2:=0,0
		if t1>=0{
			r1=int(a[t1]-'0')
		}
		if t2>=0{
			r2=int(b[t2]-'0')
		}
		sum:=r1+r2+flag
		switch sum {
		case 3:
			flag=1;result="1"+result
		case 2:
			flag=1;result="0"+result
		case 1:
			flag=0;result="1"+result
		case 0:
			flag=0;result="0"+result
		}
		t1--;t2--
	}
	if flag==1{
		result="1"+result
	}
	return result
}

func maximumGap(nums []int) int {
	sort.Ints(nums)
	var max =-1
	for i:=0;i< len(nums)-1;i++{
		xy:=abs(nums[i],nums[i+1])
		if max<xy{
			max=xy
		}
	}
	return max
}

func abs(a,b int)int{
	res:=a-b
	if res<0{
		return -res
	}
	return res
}


func titleToNumber(s string) int {

	var result =0
	for i:=0;i< len(s);i++{
		result=result*26+int(s[i]-'A'+1)
	}
	return result
}

func isPalindrome(s string) bool {
	if len(s)<1{
		return true
	}
	s1:=[]byte(s)
	for i:=0;i< len(s1);i++{
		if 'A'<=s1[i]&&s1[i]<='Z'{
			s1[i]+='a'-'A'
		}
	}
	var i,j=0, len(s1)-1
	for i<=j{
		if (s1[i]>'9'||s1[i]<'0') && (s1[i]<'a'||s1[i]>'z'){
			i++
			continue
		}
		if (s1[j]>'9'||s1[j]<'0') && (s1[j]<'a'||s1[j]>'z'){
			j--
			continue
		}
		if s1[i]!=s1[j]{
			return false
		}
		i++
		j--
	}
	return true
}
