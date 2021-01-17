package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
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
