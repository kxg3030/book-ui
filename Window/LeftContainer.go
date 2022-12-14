package Window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MenuInfo struct {
	Title string
	Intro string
	View  func(w fyne.Window) fyne.CanvasObject
}

var (
	// MenuList 所有菜单
	MenuList = map[string]MenuInfo{
		"welcome": {
			Title: "账号登陆",
			Intro: "账号登陆",
			View:  WelcomeWindow,
		},
		"fileCheck": {
			Title: "文件审核",
			Intro: "文件审核",
			View:  CheckWindowView,
		},
		"fileList": {
			Title: "文件列表",
			Intro: "文件列表",
			View:  FileListWindowView,
		},
		"fileManage": {
			Title: "文件管理",
			Intro: "文件管理",
			View:  FileManageWindow,
		},
		"other": {
			Title: "测试菜单",
			View:  OtherWindow,
		},
	}
	// MenuLayout 所有菜单组成的层级结构
	MenuLayout = map[string][]string{
		"fileManage": {"fileCheck", "fileList"},
		"":           {"welcome", "fileManage", "other"},
	}
)

type LeftContainer struct {
	MenuBox    *fyne.Container
	MenuTree   *widget.Tree
	MenuIndex  string
	LogoButton *widget.Button
	MainWindow fyne.Window
}

func NewLeftContainer(window fyne.Window) *LeftContainer {
	return &LeftContainer{
		MenuIndex:  "welcome",
		MainWindow: window,
	}
}

func (i *LeftContainer) Init() *LeftContainer {
	i.MenuTree = &widget.Tree{
		ChildUIDs: func(uid widget.TreeNodeID) (c []widget.TreeNodeID) {
			return MenuLayout[uid]
		},
		IsBranch: func(uid widget.TreeNodeID) (ok bool) {
			children, ok := MenuLayout[uid]
			return len(children) > 0 && ok
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			// 菜单初始化的时候会默认创建一个菜单
			return widget.NewLabel("默认菜单")
		},
		UpdateNode: func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
			// 在这里把默认菜单更新为预定的菜单
			menu, ok := MenuList[uid]
			if !ok {
				return
			}
			if label, ok := node.(*widget.Label); ok {
				label.SetText(menu.Title)
			}
		},
		OnSelected: func(uid widget.TreeNodeID) {
			// 当前菜单是否存在
			if _, ok := MenuList[uid]; !ok {
				return
			}
			menu, _ := MenuList[uid]
			i.MenuIndex = uid
			// 记录当前选中的菜单选项
			fyne.CurrentApp().Preferences().SetString("selectedMenu", uid)
			// 更新菜单
			application.RightBox.ContentMainBox.Objects = []fyne.CanvasObject{menu.View(i.MainWindow)}
			application.RightBox.BreadLabel.SetText(menu.Title)
			application.RightBox.ContentMainBox.Refresh()
		},
	}
	// 选中默认项
	i.MenuTree.Select(fyne.CurrentApp().Preferences().StringWithFallback(i.MenuIndex, "welcome"))
	i.LogoButton = widget.NewButtonWithIcon("", theme.ComputerIcon(), func() {
		//dialog.ShowInformation("提示", "欢迎使用超级助手", i.MainWindow)
	})
	i.MenuBox = container.NewBorder(nil, i.LogoButton, nil, nil, i.MenuTree)
	return i
}
