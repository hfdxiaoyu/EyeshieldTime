package main

import (
	"EyeshieldTime/ui"
	"EyeshieldTime/util"
	"os"
)

//主方法
func main() {
	//设置中文字体
	util.InitFront()
	//ui.InitWindow()
	ui.MainWindow();
	//取消设置单个环境变量
	os.Unsetenv("FYNE_FONT")
}
