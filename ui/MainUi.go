package ui

import (
	"EyeshieldTime/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

//主窗口
func MainWindow() {
	//创建一个app对象，使用appid创建可以推送消息到系统
	a := app.NewWithID("小绿护眼助手")
	//窗口对象
	w := a.NewWindow("小绿护眼助手")
	//设置窗口大小
	w.Resize(fyne.NewSize(200, 100))

	//时间标签
	bHour := widget.NewLabel("时：")
	bMin := widget.NewLabel("分：")
	bSecods := widget.NewLabel("秒：")

	//下拉选择框 输入时间 调用生成时间的方法
	xHour := widget.NewSelectEntry(util.NumStringBuild(24))   //时
	xMin := widget.NewSelectEntry(util.NumStringBuild(60))    //分
	xSecods := widget.NewSelectEntry(util.NumStringBuild(60)) //秒
	//设置初始值
	xHour.SetText("0")
	xMin.SetText("0")
	xSecods.SetText("0")

	//用来控制定时器
	var da = make(chan int)

	//开始按钮
	start := widget.NewButton("开始", func() {

		//TODO 这里需要进行错误判断，没点击下拉框输入框的时候里面是没有值的

		//把时间转为int类型,这里没有进行错误处理
		inxHour, _ := strconv.Atoi(xHour.Text)
		inxMin, _ := strconv.Atoi(xMin.Text)
		inxSecods, _ := strconv.Atoi(xSecods.Text)
		//开启一个协程进行计时
		go DateDecres(inxHour, inxMin, inxSecods, da)
	})

	//停止按钮
	stop := widget.NewButton("停止", func() {
		//向定时器发送一个数字，定时器收到就会调用停止的方法
		da <- 1
	})

	//布局
	//先把元素存入盒子
	//网格布局
	box := container.NewGridWithColumns(6, bHour, xHour, bMin, xMin, bSecods, xSecods)
	//HBox布局，左右堆叠
	bubox := container.NewHBox(start, stop)

	//边布局
	content := container.New(
		layout.NewBorderLayout(box, nil, nil, bubox),
		box, bubox,
	)
	//把布局加入到窗体中
	w.SetContent(content)
	w.ShowAndRun()
}

//计时器
//params: hours 时, min 分, secods 秒  da 控制计时器停止的channel
func DateDecres(hours int, min int, secods int, da chan int) {
	//time.Duration将int类型转为时间类型
	fhour := time.Duration(hours) * time.Second * 60 * 60
	fmin := time.Duration(min) * time.Second * 60
	fsecods := time.Duration(secods) * time.Second

	timer := time.NewTimer(fhour + fmin + fsecods)
	select {
	//收到消息则表示计时结束
	case <-timer.C:
		fmt.Printf("运行时间：%d 小时 %d 分钟 %d 秒\n", hours, min, secods)
		//系统弹出提示休息眼睛
		PopPushInfo(util.TimeInfoBuilder(hours, min, secods))
		break
	//收到消息就停止定时器
	case <-da:
		timer.Stop()
		break
	}

}
