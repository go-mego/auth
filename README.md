# Auth [![GoDoc](https://godoc.org/github.com/go-mego/auth?status.svg)](https://godoc.org/github.com/go-mego/auth)

Auth 套件能夠協助開發者進行最簡單的 HTTP Authorization 標頭使用者身份驗證。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)
    * [多組帳號](#多組帳號)
    * [自訂驗證](#自訂驗證)
    * [單一路由](#單一路由)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/auth
```

# 使用方式

透過 Mego 所提供的 `Use` 來將 Auth 中介軟體安插進全域路由中並開始使用。透過 `Basic` 能夠以最簡單的方式設置路由的帳號密碼。

```go
package main

import (
	"github.com/go-mego/auth"
	"github.com/go-mego/mego"
)

func main() {
	m := mego.New()
	// 所有路由都必須符合這組帳號與密碼。
	m.Use(auth.Basic("myUsername", "myPassword"))
	m.Run()
}
```

## 多組帳號

透過 `auth.Accounts` 可以傳入多組帳號與密碼，只要符合其中一組即為通過。

```go
// 所有路由都必須符合其中一組帳號與密碼。
m.Use(auth.Basic(auth.Accounts{
    "myUsername": "myPassword",
    "myAdmin":    "mySecret",
}))
```

## 自訂驗證

透過 `auth.BasicFunc`，開發者能使用自己的驗證函式來檢驗接收到的帳號與密碼是否正確。

```go
// 開發者可自訂驗證函式來確認客戶端發送的帳號與密碼是否正確，
// 這很適合與資料庫相互連接。
m.Use(auth.BasicFunc(func(username, password string) bool {
    return username == "myUsername" && password == "myPassword"
}))
```

## 單一路由

Auth 中介軟體也能夠套用到單一路由並與其他路由區隔驗證邏輯。

```go
// 限制單個路由必須符合指定的帳號與密碼。
m.Get("/", auth.Basic("myUsername", "myPassword"), func() string {
    return "哈囉，世界！"
})
```