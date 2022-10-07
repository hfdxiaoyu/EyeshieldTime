package util

import (
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"strconv"
	"strings"
	"time"
)

//初始化中文字体
func InitFront() {
	//获取系统中所有字体的列表
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//比较simkai.ttf是否在元素中
		if strings.Contains(path, "simkai.ttf") {
			//设置字体
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

//更新时间
func UpdateTime(clock *widget.Label) {
	//获取当前时间
	formatted := time.Now().Format("时间：03:04:05")
	clock.SetText(formatted)
}

//生成string类型的数字数组
//params int 需要生成的数组大小
//return []string
func NumStringBuild(num int) []string {
	var s1 = make([]string, num+1)
	for i := 0; i <= num; i++ {
		s1[i] = strconv.Itoa(i)
	}
	return s1
}

//根据传入的时间数据返回时间消息，用于推送消息的时候显示时间
//params: hours 时, min 分, secods 秒
//return 时间格式的字符串
func TimeInfoBuilder(hours int, min int, secods int) string {
	//判断返回的消息是 时还是分还是秒
	if hours > 0 && min >= 0 && secods >= 0 {
		return " " + strconv.Itoa(hours) + " 小时 " + strconv.Itoa(min) + " 分钟 " + strconv.Itoa(secods) + " 秒 "
	} else if min > 0 && secods >= 0 {
		return " " + strconv.Itoa(min) + " 分钟 " + strconv.Itoa(secods) + " 秒 "
	} else if secods > 0 {
		return " " + strconv.Itoa(secods) + " 秒 "
	}
	return " 0 小时 0 分钟 0 秒"
}
