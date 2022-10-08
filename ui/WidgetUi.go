package ui

import (
	"EyeshieldTime/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//一些其它的窗体

//护眼提醒
//timeinfo 用眼时间
func PopPushInfo(timeinfo string) {
	//消息推送到系统消息显示
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "护眼提示",
		Content: "已经用眼" + timeinfo + "啦，休息一下眼睛吧",
	})
}

//进度条显示时间倒计时
func showProgressBar(hours int, min int, secods int, bar *widget.ProgressBar, control chan int) {

	//调用更新进度条的方法
	util.UpdateProgressBar(hours, min, secods, bar, control)
}
