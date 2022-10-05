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
func InitFront()  {
	//获取系统中所有字体的列表
	fontPaths := findfont.List()
	for _,path := range fontPaths{
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//比较simkai.ttf是否在元素中
		if strings.Contains(path,"simkai.ttf") {
			//设置字体
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

//更新时间
func UpdateTime(clock *widget.Label){
	//获取当前时间
	formatted := time.Now().Format("时间：03:04:05")
	clock.SetText(formatted)
}

//生成string类型的数字数组
//params int 需要生成的数组大小
//return []string
func NumStringBuild(num int)[]string{
	var s1 = make([]string,num+1)

	for i := 0; i <= num; i++ {
		s1[i] = strconv.Itoa(i)
	}
	return s1
}


