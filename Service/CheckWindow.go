package Service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func CheckWindowView(window fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		container.NewHBox(
			widget.NewHyperlink("fyne.io", ParseURL("https://baidu.com")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", ParseURL("https://developer.fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", ParseURL("https://fyne.io/sponsor/")),
		),
		widget.NewLabel(""),
	))
}

func ParseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	return link
}
