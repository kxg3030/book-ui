package Service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FileManageWindow(window fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("文件管理默认窗口", func() {

		}),
	)
}
