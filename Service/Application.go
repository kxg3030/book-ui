package Service

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/kxg3030/book-ui/Theme"
)

//go:embed ico.ico
var iconByte []byte
var application *Application

type Application struct {
	// 应用
	app fyne.App
	// 主窗口
	Window fyne.Window
	// 左侧菜单
	LeftBox *LeftContainer
	// 右侧内容
	RightBox *RightContainer
	// 左侧菜单索引
	menuIndex string
	height    float32
	width     float32
}

func NewApp() *Application {
	application = &Application{
		height:    600,
		width:     800,
		menuIndex: "default",
	}
	return application
}

func (i *Application) Run() {
	i.height = 600
	i.width = 800
	i.app = app.NewWithID("com.book.ui")
	// 保存全局对象
	i.SaveApp()
	// 设置支持中文
	i.SupportZh()
	// 设置主窗口
	i.MainWindow()
}

func (i *Application) SaveApp() {

}

func (i *Application) SupportZh() {
	i.app.Settings().SetTheme(&Theme.Theme{})
}

func (i *Application) MainWindow() {
	i.Window = i.app.NewWindow("book-ui")
	// 设置托盘图标
	i.MakeTray()
	// 设置应用图标
	i.Iconic()
	// 设置右侧容器
	i.RightBoxSet()
	// 左侧菜单容器
	i.LeftMenuSet()
	// 显示主窗口
	i.MasterWindow()
}

func (i *Application) MasterWindow() {
	split := container.NewHSplit(i.LeftBox.MenuBox, i.RightBox.ContentBox)
	split.Offset = 0.25
	i.Window.Resize(fyne.NewSize(i.width, i.height))
	i.Window.SetContent(split)
	i.Window.SetMaster()
	i.Window.ShowAndRun()
}

// LeftMenuSet 左侧菜单
func (i *Application) LeftMenuSet() {
	i.app.Preferences().SetString("selectedMenu", "check")
	i.LeftBox = NewLeftContainer(i.Window).Init()
}

// RightBoxSet 右侧内容
func (i *Application) RightBoxSet() {
	i.RightBox = NewRightContainer()
	i.RightBox.Init()
}

// MakeTray 托盘图标
func (i *Application) MakeTray() {
	//i.Window.SetCloseIntercept(func() {
	//	i.Window.Hide()
	//})
	if desk, ok := i.app.(desktop.App); ok {
		menu := fyne.NewMenu("")
		menu.Items = make([]*fyne.MenuItem, 0)
		// 创建子菜单
		item := fyne.NewMenuItem("显示", func() {
			i.Window.Show()
			menu.Refresh()
		})
		menu.Items = append(menu.Items, item)
		desk.SetSystemTrayMenu(menu)
	}
}

// Iconic 系统图标
func (i *Application) Iconic() {
	fyne.CurrentApp().SetIcon(&fyne.StaticResource{
		StaticName:    "ico.ico",
		StaticContent: iconByte,
	})
}
