package Window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func WelcomeWindow(window fyne.Window) fyne.CanvasObject {
	username := widget.NewEntry()
	username.SetPlaceHolder("请输入账号")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("请输入密码")
	loginForm := widget.NewForm(widget.NewFormItem("账号", username), widget.NewFormItem("密码", password))
	loginButton := widget.NewButtonWithIcon("登陆", theme.LoginIcon(), func() {
		if len(username.Text) <= 0 || len(password.Text) <= 0 {
			dialog.ShowInformation("错误", "账号或密码不能为空", window)
			return
		}
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "提示",
			Content: "登陆成功",
		})
	})
	notice := widget.NewRichTextFromMarkdown(`
- 说明：仅用于验证，不会收集任何信息
`)

	logo := canvas.NewImageFromResource(theme.MediaMusicIcon())
	logo.FillMode = canvas.ImageFillStretch
	logo.SetMinSize(fyne.NewSize(0, 300))
	return container.NewVBox(loginForm, loginButton, notice, logo)
}
