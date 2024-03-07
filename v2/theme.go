package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	fyne.Theme
}

func MyTheme() fyne.Theme {
	return myTheme{theme.DefaultTheme()}
}

// return bundled font resource
func (t myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return t.Theme.Font(s)
	}
	if s.Bold {
		if s.Italic {
			return t.Theme.Font(s)
		}
		return resourceMplus1cBoldTtf
	}
	if s.Italic {
		return t.Theme.Font(s)
	}
	return resourceMplus1cRegularTtf
}
