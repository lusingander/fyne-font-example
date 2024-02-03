
fyne-font-example
====

[English](./README.md) / 日本語

## About

[Fyne](https://fyne.io) で日本語フォントを利用するサンプルアプリケーションです。

> [!IMPORTANT]
> これは Fyne v2.x についての説明です. Fyne v1.x 以前のバージョンについて知りたい場合は [v1](./v1) 以下を参照してください.

<img src="./resource/image-v2.png" width=300>

> [!TIP]
> 🎨 GUI でテーマを作成したい場合は [fyne-theme-generator](https://github.com/lusingander/fyne-theme-generator) も試してみてください.

## Summary

#### 0. `fyne` コマンドをインストール

```
$ go get fyne.io/fyne/v2/cmd/fyne

$ fyne
Usage: fyne [command] [parameters], where command is one of:
...
```

#### 1. フォントファイルを用意して `fyne bundle` コマンドを実行

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

詳細は [./v2/bundle.go](./v2/bundle.go) を参照してください.

> [!WARNING]
> ファイルサイズがかなり大きくなります

#### 2. カスタムテーマを作成しフォントリソースを読み込む

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

詳細は [./v2/theme.go](./v2/theme.go) を参照してください.

#### 3. カスタムテーマを読み込む

```go
...
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
...
```

詳細は [./v2/main.go](./v2/main.go) を参照してください.


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

----

サンプルのフォントとして [M+ FONTS](http://mplus-fonts.osdn.jp/) を使用しています.

http://mplus-fonts.osdn.jp/
