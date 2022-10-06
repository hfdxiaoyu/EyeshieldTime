package ui

import (
	"fyne.io/fyne/v2"
)

//一些其它的窗体

//相同显示消息提示
func PopPushInfo(timeinfo string) {
	//消息推送到系统消息显示
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "护眼提示",
		Content: "已经用眼" + timeinfo + "啦，休息一下眼睛吧",
	})
}
