package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)


type blankSt struct {
	a int
	_ string
}
func TestChange(t *testing.T){
	fmt.Println(DealLength("67+036.5","67+627.5"))
}
func DealLength(str1,str2 string)(string,error){
	s1:=strings.ReplaceAll(str1,"+","")
	f1,err:=strconv.ParseFloat(s1,10)
	if err!=nil{
		return "", err
	}
	s2:=strings.ReplaceAll(str2,"+","")
	f2,err:=strconv.ParseFloat(s2,10)
	if err!=nil{
		return "", err
	}
	return strconv.FormatFloat(math.Abs(f1-f2),'f',2,64),nil
}
func DealExcelTime(str *string) (err error) {
	if str == nil || *str == ""{
		return
	}
	atto, err := strconv.Atoi(*str)
	if err != nil {
		*str=""
		return
	}
	excelTime := time.Date(1900, time.January, -1, 0, 0, 0, 0, time.Local)
	*str = excelTime.Add(time.Duration(int(time.Hour) * 24 * atto)).Format("2006-01-02")
	return nil
}