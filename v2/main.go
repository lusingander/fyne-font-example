package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(MyTheme())
	w := a.NewWindow("font")
	w.Resize(fyne.NewSize(300, 200))
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			widget.NewLabel("こんにちは、ファイン"),
			widget.NewLabel("これは日本語のラベルです"),
			widget.NewButton("これはボタンです", func() {
				dialog.ShowInformation("確認", "これはダイアログです", w)
			}),
			layout.NewSpacer(),
		),
	)
	w.ShowAndRun()
}
