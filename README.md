# Utility library for Golang

## Overview

Golangの共通ライブラリとして処理をまとめてあります。  
現在は以下のようなことができます。

1. 設定ファイルをJSONより読み込み、構造体に保持
2. ログ出力をレベル分けして出力
3. メール送信

## Requirements

- Golang 1.8+

## Setup

以下コマンドで`$GOPATH`にインストールされます。

```sh
$ go get -u github.com/gotoeveryone/golib
```

## Run

`config.json.example`を参考に、任意ディレクトリに「config.json」を作成してください。  
※値は実際に利用するサービスの接続情報を設定すること。  

以下コマンドで実行時、設定ファイルの格納ディレクトリを示す`--conf`オプションで対象ディレクトリを指定してください。  
未指定の場合は実行ファイルと同じディレクトリを参照します。  

```sh
$ go run <my_program> --conf=/path/to/
```

## Test

```sh
$ cd <this_directory>
$ go test
```
