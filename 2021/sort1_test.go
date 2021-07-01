package _021

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type (
	Tunnel struct {
		TunnelRockBasic
		TunnelRockSubExt
	}
	//基本信息;
	TunnelRockBasic struct {
		TunnelUnitProjIdUnion     int     `json:"unit_proj_id_union"      comment:"隧道单位工程id关联"`
		TunnelRockSubUnitName     string  `json:"rock_sub_unit_name"      comment:"子单位工程"`
		TunnelRockPilePrefix      string  `json:"rock_pile_prefix"        comment:"围岩等级桩号前缀"`
		TunnelRockPileStartString string  `json:"rock_pile_start_string"  comment:"围岩等级起始桩号-a+b"`
		TunnelRockPileEndString   string  `json:"rock_pile_end_string"    comment:"围岩等级终止桩号-a+b"`
		TunnelRockPileStart       float64 `json:"rock_pile_start"         comment:"围岩等级起始桩号"`
		TunnelRockPileEnd         float64 `json:"rock_pile_end"           comment:"围岩等级终止桩号"`
		TunnelRockPileLength      float64 `json:"rock_pile_length"        comment:"长度"`
		TunnelAssist              string  `json:"tunnel_assist"           comment:"辅助工程措施"`
		TunnelRockDesc            string  `json:"rock_desc"               comment:"备注"`
		TunnelRockT
		TunnelLiningTypeT
	}
	TunnelRockT struct {
		TunnelRock string `json:"rock"   comment:"围岩等级"`
	}
	TunnelLiningTypeT struct {
		TunnelLiningType string `json:"lining_type"  comment:"衬砌类型"`
	}
	TunnelRockSubExt struct {
		TunnelRockSubAccountIdUn int    `json:"rock_sub_account_id_un"  comment:"-服务器自动填写-上传人员id"`
		TunnelRockSubTime        string `json:"rock_sub_time"           comment:"-服务器自动填写-上传时间 format:2006-01-02 15:04:05"`
		TunnelRockSubTimeInt     int64  `json:"rock_sub_time_int"       comment:"-服务器自动填写-上传时间int format:20060102150405"`
	}
)

//排序函数
type By func(p1 *Tunnel,p2 *Tunnel)bool

//实现函数
func (by By)Sort(data []Tunnel){
	s:=&TunnelSort{
		Data: data,
		sortFunc: by,
	}
	sort.Sort(s)
}

type TunnelSort struct {
	Data []Tunnel
	sortFunc func(p1 *Tunnel,p2 *Tunnel)bool
}
//实现排序方法
func (this *TunnelSort)Len()int{
	return len(this.Data)
}
func (this *TunnelSort)Swap(i,j int){
	this.Data[i],this.Data[j]=this.Data[j],this.Data[i]
}
func (this *TunnelSort)Less(i,j int)bool{
	return this.sortFunc(&this.Data[i],&this.Data[j])
}

func Test_t1(t *testing.T) {

	var tunnelArray =[]Tunnel{
		{
			TunnelRockBasic:TunnelRockBasic{
				TunnelRockSubUnitName:"金宝石大桥左线",
				TunnelRockPileStart:1000,
				TunnelRockPileEnd:2000,
				TunnelRockPileLength:1000,
				TunnelRockDesc:"世界上最大的桥梁",
			},
			TunnelRockSubExt:TunnelRockSubExt{
				TunnelRockSubTime:time.Now().Format("2006-01-02"),
				TunnelRockSubTimeInt:time.Now().Unix(),
			},

		},
		{
			TunnelRockBasic:TunnelRockBasic{
				TunnelRockSubUnitName:"郎寨互通",
				TunnelRockPileStart:1000,
				TunnelRockPileEnd:3000,
				TunnelRockPileLength:2000,
				TunnelRockDesc:"七月七日长生殿,月下无人私语时",
			},
			TunnelRockSubExt:TunnelRockSubExt{
				TunnelRockSubTime:time.Now().Format("2006-01-02"),
				TunnelRockSubTimeInt:time.Now().Unix(),
			},

		},
		{
			TunnelRockBasic:TunnelRockBasic{
				TunnelRockSubUnitName:"粟木寨互通",
				TunnelRockPileStart:500,
				TunnelRockPileEnd:600,
				TunnelRockPileLength:100,
				TunnelRockDesc:"千呼万唤始出来,犹抱琵琶半遮面",
			},
			TunnelRockSubExt:TunnelRockSubExt{
				TunnelRockSubTime:time.Now().Format("2006-01-02"),
				TunnelRockSubTimeInt:time.Now().Unix(),
			},

		},
	}

	//按照名字排序的策略
	name:=func(p1 ,p2 *Tunnel)bool{
		return p1.TunnelRockSubUnitName<p2.TunnelRockSubUnitName
	}
	//按照长度排序的策略
	length:=func(p1 ,p2 *Tunnel)bool{
		return p1.TunnelRockPileLength<p2.TunnelRockPileLength
	}
	//按照创建时间排序
	create:=func(p1 ,p2 *Tunnel)bool{
		return p1.TunnelRockSubTimeInt<p2.TunnelRockSubTimeInt
	}
	By(name).Sort(tunnelArray)
	fmt.Println(tunnelArray)
	By(length).Sort(tunnelArray)
	fmt.Println(tunnelArray)
	By(create).Sort(tunnelArray)
	fmt.Println(tunnelArray)
}
