package Service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RightContainer struct {
	Title          string
	BreadLabel     *widget.Label
	ContentBox     *fyne.Container
	ContentMainBox *fyne.Container
}

func NewRightContainer() *RightContainer {
	return &RightContainer{
		Title: "默认标题",
	}
}

func (i *RightContainer) Init() *RightContainer {
	i.BreadLabel = widget.NewLabel(i.Title)
	i.ContentMainBox = container.NewMax()
	i.ContentBox = container.NewBorder(
		container.NewVBox(i.BreadLabel, widget.NewSeparator()),
		nil, nil, nil, i.ContentMainBox,
	)
	return i
}
