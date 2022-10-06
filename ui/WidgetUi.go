package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

//一些其它的窗体

//弹窗
func PopUpWindow(window fyne.Window) {

	//设置弹窗大小
	window.Resize(fyne.NewSize(230, 150))
	confirm := dialog.NewConfirm("护眼提示", "休息一下眼睛吧", func(b bool) {
		if b {
			window.Hide()
		} else {
			window.Hide()
		}
	}, window)

	confirm.Show()
	window.Show()
}
