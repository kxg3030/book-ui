package Window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kxg3030/book-ui/Service"
	"golang.org/x/image/colornames"
)

func WelcomeWindow(window fyne.Window) fyne.CanvasObject {
	username := widget.NewEntry()
	username.SetPlaceHolder("请输入账号")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("请输入密码")
	loginForm := widget.NewForm(widget.NewFormItem("账号", username), widget.NewFormItem("密码", password))

	statusLabel := canvas.NewText("未登陆", colornames.Orange)
	// 如果要居中,最外层只能border容器,内部包裹一个max容器(max容器才有高度)
	statusBox := container.NewMax(container.NewCenter(statusLabel))
	// 账号登陆
	loginButton := widget.NewButtonWithIcon("登陆", theme.LoginIcon(), func() {
		if len(username.Text) <= 0 || len(password.Text) <= 0 {
			dialog.ShowInformation("错误", "账号或密码不能为空", window)
			return
		}
		application.LeftBox.LogoButton.SetText(username.Text)
		statusLabel.Text = "已登陆"
		statusLabel.Color = colornames.Green
		statusLabel.Refresh()
		success := Service.NewBookService(username.Text, password.Text).Login()
		notice := &fyne.Notification{Title: "提示", Content: "账号或密码错误"}
		if success {
			notice.Content = "登陆成功"
		}
		fyne.CurrentApp().SendNotification(notice)
	})
	// 退出登陆
	cancelButton := widget.NewButtonWithIcon("退出", theme.LogoutIcon(), func() {
		username.SetText("")
		password.SetText("")
		application.LeftBox.LogoButton.SetText("")
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "提示",
			Content: "退出登陆",
		})
	})
	// 登陆按钮容器
	grid := container.NewGridWithColumns(2, loginButton, cancelButton)

	return container.NewBorder(container.NewVBox(loginForm, grid), nil, nil, nil, statusBox)
}
