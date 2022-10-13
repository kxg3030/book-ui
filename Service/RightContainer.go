package Service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RightContainer struct {
	Title           string
	BreadLabel      *widget.Label
	ContentBox      *fyne.Container
	ContentMainBox  *fyne.Container
	ContentMainNext *fyne.Container
}

func NewRightContainer() *RightContainer {
	return &RightContainer{
		Title: "默认标题",
	}
}

func (i *RightContainer) Init() *RightContainer {
	i.BreadLabel = widget.NewLabel(i.Title)
	i.ContentMainNext = container.NewMax()
	i.ContentMainBox = container.NewBorder(nil, nil, nil, nil, i.ContentMainNext)
	i.ContentBox = container.NewVBox(i.BreadLabel, widget.NewSeparator(), i.ContentMainBox)
	return i
}
