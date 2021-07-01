package tw

import (
	"fmt"
	"testing"
)

func Test_distribute(t *testing.T){

	fmt.Println(countCharacters([]string{"cat","bt","hat","tree"},"atach"))
}
func distributeCandies2(candies int, num_people int) []int {
	var result =make([]int,num_people)
	i:=0
	for candies>0{
		result[i%num_people]+=minI(i+1,candies)
		candies-=minI(i+1,candies)
		i++
	}
	return result
}
func minI(a,b int)int{
	if a<b{
		return a
	}
	return b
}


func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var result =make([]int, len(queries))
	sum:=0
	for i:=0;i< len(A);i++{
		if A[i]&1==0{
			sum+=A[i]
		}
	}
	for i:=0;i< len(queries);i++{
		index:=queries[i][1]
		if A[index]&1==0{
			sum-=A[index]
		}
		A[index]+=queries[i][0]
		if A[index]&1==0{
			sum+=A[index]
		}
		result[i]=sum
	}
	return result
}

func s2int(str string)int{
	var result =0
	for i:=0;i< len(str);i++{
		result=result*10+int(str[i]-'0')
	}
	return result
}

//二月份的天数
func month(str string)int{
	year := s2int(str)
	if year%4==0&&year%100!=0{
		return 29
	}else{
		return 28
	}
}

func dayOfYear(date string) int {
	var space =0
	var cons =[]int{31,month(date[0:4]),31,30,31,30,31,31,30,31,30}
	index:=s2int(date[5:7])-1
	for i:=0;i<index;i++{
		space+=cons[i]
	}
	return space+s2int(date[8:])
}


///////////////////////////l////////////////////////////
func countCharacters(words []string, chars string) int {

	var set =make([]int,26)
	var temp=make([]int,26)
	var result =0
	for i:=0;i< len(chars);i++{
		set[chars[i]-'a']++
	}
	for i:=0;i< len(words);i++{
		copy(temp,set)
		j:=0
		for ;j< len(words[i]);j++{
			temp[words[i][j]-'a']--
			if temp[words[i][j]-'a']<0{
				break
			}
		}
		if j== len(words[i]){
			result+= len(words[i])
		}
	}
	return result
}
