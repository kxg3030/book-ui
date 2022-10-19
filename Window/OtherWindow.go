package Window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func OtherWindow(window fyne.Window) fyne.CanvasObject {
	statusLabel := widget.NewLabel("测试")
	return container.NewCenter(statusLabel)
}
