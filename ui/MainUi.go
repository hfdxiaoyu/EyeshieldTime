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

func MainWindow() {
	a := app.New()
	w := a.NewWindow("护眼助手")
	w1 := a.NewWindow("护眼提示")
	//设置窗口大小
	w.Resize(fyne.NewSize(100, 100))

	//时间标签
	bHour := widget.NewLabel("时：")
	bMin := widget.NewLabel("分：")
	bSecods := widget.NewLabel("秒：")

	//下拉选择框 输入时间
	xHour := widget.NewSelectEntry(util.NumStringBuild(24))   //时
	xMin := widget.NewSelectEntry(util.NumStringBuild(60))    //分
	xSecods := widget.NewSelectEntry(util.NumStringBuild(60)) //秒

	//用来控制定时器
	var da = make(chan int)

	//开始按钮
	start := widget.NewButton("开始", func() {
		//把时间转为int类型
		inxHour, _ := strconv.Atoi(xHour.Text)
		inxMin, _ := strconv.Atoi(xMin.Text)
		inxSecods, _ := strconv.Atoi(xSecods.Text)

		go DateDecres(inxHour, inxMin, inxSecods, da, w1)
	})

	//停止按钮
	stop := widget.NewButton("停止", func() {
		fmt.Println("停止")
		da <- 1
	})

	//布局
	//先把元素存入盒子
	//网格布局
	box := container.NewGridWithColumns(6, bHour, xHour, bMin, xMin, bSecods, xSecods)
	bubox := container.NewHBox(start, stop)

	//边布局
	content := container.New(
		layout.NewBorderLayout(box, nil, nil, bubox),
		box, bubox,
	)
	w.SetContent(content)
	w.ShowAndRun()
}

//计时器
func DateDecres(hours int, min int, secods int, da chan int, wi fyne.Window) {
	//time.Duration将int类型转为时间类型
	fhour := time.Duration(hours) * time.Second * 60 * 60
	fmin := time.Duration(min) * time.Second * 60
	fsecods := time.Duration(secods) * time.Second

	timer := time.NewTimer(fhour + fmin + fsecods)
	select {
	case <-timer.C:
		fmt.Printf("运行时间：%d 小时 %d 分钟 %d 秒\n", hours, min, secods)
		PopUpWindow(wi)
		break
	case <-da:
		timer.Stop()
		break
	}

}
