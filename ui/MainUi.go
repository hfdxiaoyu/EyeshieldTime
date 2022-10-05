package ui

import (
	"EyeshieldTime/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainWindow(){
	a := app.New()
	w := a.NewWindow("护眼助手")
	//设置窗口大小
	w.Resize(fyne.NewSize(100,100))

	//时间标签
	bHour := widget.NewLabel("时：")
	bMin := widget.NewLabel("分：")
	bSecods := widget.NewLabel("秒：")


	//数据绑定
	//bdHour := binding.NewString()
	//bdMin := binding.NewString()
	//bdSecods := binding.NewString()

	//输入框 输入时间
	//hour := widget.NewEntryWithData(bdHour) //时
	//min := widget.NewEntryWithData(bdMin) //分
	//secods := widget.NewEntryWithData(bdSecods) //秒

	//下拉选择框 输入时间
	xHour := widget.NewSelectEntry(util.NumStringBuild(24)) //时
	xMin := widget.NewSelectEntry(util.NumStringBuild(60)) //时
	xSecods := widget.NewSelectEntry(util.NumStringBuild(60)) //时

	//xmin := widget.NewEntryWithData(bdMin) //分
	//xsecods := widget.NewEntryWithData(bdSecods) //秒


	//开始按钮
	start := widget.NewButton("开始", func() {
		dialog.NewConfirm("提示","护眼提示", func(b bool) {

		},w)
	})

	//停止按钮
	stop := widget.NewButton("停止", func() {

	})

	//布局
	//先把元素存入盒子
	//网格布局
	box := container.NewGridWithColumns(6,bHour,xHour,bMin,xMin,bSecods,xSecods)
	bubox := container.NewHBox(start,stop)

	//边布局
	content := container.New(
		layout.NewBorderLayout(box,nil,nil,bubox),
		box,bubox,
		)
	w.SetContent(content)
	w.ShowAndRun()
}

