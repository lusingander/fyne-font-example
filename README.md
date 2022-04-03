fyne-font-example
====

English / [日本語](./README_ja.md)

## About

Sample application that uses different fonts in [Fyne](https://fyne.io).

> This is a description for Fyne v2.x. If you want to know for Fyne v1.x or earlier, please refer to the [v1](./v1) directory.

<img src="./resource/image-v2.png" width=300>

> 🎨 Prefer to work with the GUI? Try the [fyne-theme-generator](https://github.com/lusingander/fyne-theme-generator)!

## Summary

#### 0. Install `fyne` command

```
$ go get fyne.io/fyne/v2/cmd/fyne

$ fyne
Usage: fyne [command] [parameters], where command is one of:
...
```

#### 1. Prepare the font file and execute `fyne bundle` command

```
$ fyne bundle mplus-1c-regular.ttf > bundle.go

$ head -n 9 bundle.go
// auto-generated

package main

import "fyne.io/fyne/v2"

var resourceMplus1cRegularTtf = &fyne.StaticResource{
	StaticName: "mplus-1c-regular.ttf",
	StaticContent: []byte{
```

See [./v2/bundle.go](./v2/bundle.go).

> Warning: the file size is very large

#### 2. Create the custom theme and load font resources

```go
type myTheme struct{}

func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceMplus1cBoldTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceMplus1cRegularTtf
}
...
```

See [./v2/theme.go](./v2/theme.go).

#### 3. Load the custom theme

```go
...
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
...
```

See [./v2/main.go](./v2/main.go).

## A little more details

`bundle.go` is generated using [fyne command](https://github.com/fyne-io/fyne/tree/master/cmd/fyne).

```
$ fyne bundle mplus-1c-regular.ttf > bundle.go
$ fyne bundle -append mplus-1c-bold.ttf >> bundle.go
```

See the Blog below for more information. (Japanese)

- [About the `fyne` command](https://lusingander.netlify.app/posts/200613-fyne-resourece/)
- [About fonts](https://lusingander.netlify.app/posts/200614-fyne-font/)

An official tutorial has also been added on resource bundling.

- [Bundling resources | Develop using Fyne](https://developer.fyne.io/tutorial/bundle)

----

[M+ FONTS](http://mplus-fonts.osdn.jp/) is included and used as a sample font file.

http://mplus-fonts.osdn.jp/
