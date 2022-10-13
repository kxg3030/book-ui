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

type MenuInfo struct {
	Title      string
	Intro      string
	SupportWeb bool
	View       func(w fyne.Window) fyne.CanvasObject
}

type Application struct {
	// 应用
	app fyne.App
	// 主窗口
	window fyne.Window
	// 左侧菜单
	left *fyne.Container
	// 右侧内容
	rightBox *RightContainer
	// 左侧菜单索引
	menuIndex string
	height    float32
	width     float32
}

func NewApp() *Application {
	return &Application{
		height:    600,
		width:     800,
		menuIndex: "default",
	}
}

func (i *Application) Run() {
	i.height = 600
	i.width = 800
	i.app = app.NewWithID("com.book.ui")
	// 设置支持中文
	i.SupportZh()
	// 设置主窗口
	i.MainWindow()
}

func (i *Application) Setting() {

}

func (i *Application) SupportZh() {
	i.app.Settings().SetTheme(&Theme.Theme{})
}

func (i *Application) MainWindow() {
	// 设置托盘图标
	i.MakeTray()
	// 设置应用图标
	i.Iconic()
	// 设置右侧容器
	i.RightBox()
	// 左侧菜单容器
	i.LeftMenu()
	// 显示主窗口
	i.MasterWindow()
}

func (i *Application) MasterWindow() {
	split := container.NewHSplit(i.left, i.rightBox.ContentBox)
	split.Offset = 0.2

	i.window = i.app.NewWindow("book-ui")
	i.window.Resize(fyne.NewSize(i.width, i.height))
	i.window.SetContent(split)
	i.window.SetMaster()
	i.window.ShowAndRun()
}

// LeftMenu 左侧菜单
func (i *Application) LeftMenu() {
	// 默认菜单选项
	i.app.Preferences().SetString("selectedMenu", "check")
	i.left = NewLeftContainer(i.rightBox, i.window).Init().MenuBox
}

// RightBox 右侧内容
func (i *Application) RightBox() {
	i.rightBox = NewRightContainer()
	i.rightBox.Init()
}

// MakeTray 托盘图标
func (i *Application) MakeTray() {
	if desk, ok := i.app.(desktop.App); ok {
		menu := fyne.NewMenu("托盘菜单")
		menu.Items = make([]*fyne.MenuItem, 0)
		// 创建子菜单
		item := fyne.NewMenuItem("测试菜单", func() {
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
