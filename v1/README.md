fyne-font-example
====


[Fyne](https://fyne.io) で日本語フォントを利用するサンプルアプリケーションです。

Sample application that uses different fonts in [Fyne](https://fyne.io).

<img src="./resource/image.png" width=300>

> If you prefer to work with the GUI, see [fyne-theme-generator](https://github.com/lusingander/fyne-theme-generator).

## Summary

#### 0. `fyne` コマンドをインストール / Install `fyne` command

```
$ go get fyne.io/fyne/cmd/fyne

$ fyne
Usage: fyne [command] [parameters], where command is one of:
...
```

#### 1. フォントファイルを用意して `fyne bundle` コマンドを実行 / Prepare the font file and execute `fyne bundle` command

```
$ fyne bundle mplus-1c-regular.ttf > bundle.go

$ head -n 9 bundle.go
// auto-generated

package main

import "fyne.io/fyne"

var resourceMplus1cRegularTtf = &fyne.StaticResource{
	StaticName: "mplus-1c-regular.ttf",
	StaticContent: []byte{
```

See [bundle.go](./bundle.go).

> Warning: the file size is very large

#### 2. カスタムテーマを作成しフォントリソースを読み込む / Create the custom theme and load font resources

```go
type myTheme struct{}

func (myTheme) TextFont() fyne.Resource     { return resourceMplus1cRegularTtf }
...
```

See [theme.go](./theme.go).

#### 3. カスタムテーマを読み込む / Load the custom theme

```go
...
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
...
```

See [main.go](./main.go).


## もう少し詳しく

`bundle.go` は [fyne command](https://github.com/fyne-io/fyne/tree/master/cmd/fyne) を利用して生成しています.

```
$ fyne bundle mplus-1c-regular.ttf > bundle.go
$ fyne bundle -append mplus-1c-bold.ttf >> bundle.go
```

詳細については以下の記事に記載しています.

- [`fyne` コマンドで各種リソースをバンドルする方法](https://lusingander.netlify.app/posts/200613-fyne-resourece/)
- [Fyne で日本語を扱う](https://lusingander.netlify.app/posts/200614-fyne-font/)

公式のチュートリアルにもリソースのバンドルについて追記されました.

- [Bundling resources | Develop using Fyne](https://developer.fyne.io/tutorial/bundle)


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
