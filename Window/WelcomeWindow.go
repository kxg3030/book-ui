package Window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WelcomeWindow(window fyne.Window) fyne.CanvasObject {
	rich := widget.NewRichTextFromMarkdown(`
## RichText Heading

## A Sub Heading

---

* Item1 in _three_ segments
* Item2
* Item3

Normal **Bold** *Italic* [Link](https://fyne.io/) and some ` + "`Code`" + `.
This styled row should also wrap as expected, but only *when required*.

> An interesting quote here, most likely sharing some very interesting wisdom.`)
	rich.Scroll = container.ScrollHorizontalOnly
	return container.NewVBox(container.NewGridWithRows(1, rich))
}
