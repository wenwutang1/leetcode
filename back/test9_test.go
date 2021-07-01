package back

import (
	"fmt"
	"testing"
)

//电子游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。
//给定一个字符串ring，表示刻在外环上的编码；给定另一个字符串key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
//最初，ring的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使key的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完key中的所有字符。
//旋转ring拼出 key 字符key[i]的阶段中：
//您可以将ring顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串ring的一个字符与 12:00 方向对齐，并且这个字符必须等于字符key[i] 。
//如果字符key[i]已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作1 步。按完之后，您可以开始拼写key的下一个字符（下一阶段）, 直至完成所有拼写。

func Test_F(t *testing.T) {
	findRotateSteps("godding", "gd")
}

func findRotateSteps(ring string, key string) int {

	var min = 0
	backRotate(ring, key, 0, 0, 0, &min)
	fmt.Println(min)
	return min
}

//ring 环
//key 目标字符
//sum 找到了的字符数量
//direct 旋转的方向
//road 当前的步数
func backRotate(ring string, key string, sum int, direct int, road int, min *int) {
	if sum == len(key) {
		if *min > road {
			*min = road
		}
		return
	}
	if direct == len(ring) {
		direct = 0
	}
	if direct < 0 {
		direct = len(ring) - 1
	}
	for i := direct; i < len(ring); i++ {
		road++
		//找到了一个字符
		if ring[i] == ring[sum] {
			sum++
		}
		//顺时针
		backRotate(ring, key, sum, i+1, road, min)
		//逆时针
		backRotate(ring, key, sum, i-1, road, min)
	}
}
