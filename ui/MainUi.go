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

	//新建一个进度条
	bar := widget.NewProgressBar()

	//用来控制定时器
	var da = make(chan int)

	//开始按钮
	start := widget.NewButton("开始", func() {

		//把时间转为int类型,这里没有进行错误处理
		inxHour, _ := strconv.Atoi(xHour.Text)
		inxMin, _ := strconv.Atoi(xMin.Text)
		inxSecods, _ := strconv.Atoi(xSecods.Text)

		//开启一个协程进行计时
		go DateDecres(inxHour, inxMin, inxSecods, da, bar)
	})

	//停止按钮
	stop := widget.NewButton("停止", func() {
		//向定时器发送一个数字，定时器收到就会调用停止的方法
		da <- -1
	})

	//重启按钮
	rest := widget.NewButton("重启", func() {
		da <- 1
	})

	//布局
	//先把元素存入盒子
	//网格布局
	box := container.NewGridWithColumns(6, bHour, xHour, bMin, xMin, bSecods, xSecods)
	//HBox布局，左右堆叠
	bubox := container.NewHBox(start, rest, stop)

	//边布局
	content := container.New(
		layout.NewBorderLayout(box, nil, bar, bubox),
		box, layout.NewSpacer(), bar, bubox,
	)
	//把布局加入到窗体中
	w.SetContent(content)
	w.ShowAndRun()
}

//计时器
//params: hours 时, min 分, secods 秒  da 控制计时器状态的channel
func DateDecres(hours int, min int, secods int, da chan int, bar1 *widget.ProgressBar) {
	//time.Duration将int类型转为时间类型
	fhour := time.Duration(hours) * time.Second * 60 * 60
	fmin := time.Duration(min) * time.Second * 60
	fsecods := time.Duration(secods) * time.Second
	//把参数传入定时器
	timer := time.NewTimer(fhour + fmin + fsecods)
	//控制进度条的channel
	var controlbar = make(chan int)
	//进度条
	showProgressBar(hours, min, secods, bar1, controlbar)
	//存储运行时间
	runtime := time.Second
	//点击停止时把已运行时间存入此变量
	stopRunTime := time.Second * 0
	//时间自动记录的协程
	go func() {
		//每秒运行一次
		for range time.Tick(time.Second) {
			runtime += time.Second
			//运行完退出循环
			if runtime > fhour+fmin+fsecods {
				break
			}
		}
	}()

	//开启协程循环监听
	go func() {
		//循环监听
		for {
			select {
			//收到消息则表示计时结束
			case <-timer.C:
				fmt.Printf("运行时间：%d 小时 %d 分钟 %d 秒\n", hours, min, secods)
				//系统弹出提示休息眼睛
				PopPushInfo(util.TimeInfoBuilder(hours, min, secods))
				break
			//根据收到的消息判断要对定时器进行的操作
			case c := <-da:
				//如果收到的是 -1 则停止定时器
				//对状态进行判断
				if c == -1 {
					//发送消息停止进度条
					controlbar <- -1
					stopRunTime = runtime
					timer.Stop()
					//如果收到的是1则重启定时器
				} else if c == 1 {
					//发送消息启动进度条
					controlbar <- 1
					timer.Reset((fhour + fmin + fsecods) - stopRunTime)
				}

			}
			//运行结束停止监听
			if runtime > fhour+fmin+fsecods {
				break
			}
		}
	}()

}
