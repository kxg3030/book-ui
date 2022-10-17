package Service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FileListWindowView(window fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("测试", func() {

		}),
	)
}
