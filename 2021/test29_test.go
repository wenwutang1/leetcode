package _021

import (
	"fmt"
	"testing"
	"time"
)

func Test29(t *testing.T) {
	//now:=time.Now()
	//xiaban,err:=time.ParseInLocation("2006-01-02 15:04:05","2021-04-23 18:00:00",time.Local)
	//if err!=nil{
	//	return
	//}
	//nowUnix := now.Unix()
	//
	//xiaban.Add(-time.Duration(10+10)*time.Second*60)
	//
	//xiaBanUnix := xiaban.Unix()
	//
	//fmt.Println(nowUnix,xiaBanUnix)

	location, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-04-23 18:00:00", time.Local)
	location=location.Add(time.Minute)
	fmt.Println(location.Format("2006-01-02 15:04:05"))
}

