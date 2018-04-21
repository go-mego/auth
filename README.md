# Auth [![GoDoc](https://godoc.org/github.com/go-mego/auth?status.svg)](https://godoc.org/github.com/go-mego/auth)

Auth 套件能夠協助開發者進行最簡單的 HTTP Authorization 標頭使用者身份驗證。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)
    * [多組帳號](#多組帳號)
    * [自訂驗證](#自訂驗證)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/auth
```

# 使用方式

透過 `auth.New` 能夠以最簡單的方式設置路由的帳號密碼，將其傳入 Mego 所提供的 `Use` 來將 Auth 中介軟體安插進全域中介軟體並開始使用。

```go
package main

import (
	"github.com/go-mego/auth"
	"github.com/go-mego/mego"
)

func main() {
	m := mego.New()
	// 所有路由都必須符合這組帳號與密碼。
	m.Use(auth.New("myUsername", "myPassword"))
	m.Run()
}
```

Auth 中介軟體也能夠套用到單一路由並與其他路由區隔驗證邏輯。

```go
func main() {
	m := mego.New()
	// 限制單個路由必須符合指定的帳號與密碼。
	m.Get("/", auth.New("myUsername", "myPassword"), func() string {
		return "哈囉，世界！"
	})
	m.Run()
}
```

## 多組帳號

透過 `auth.NewAccounts` 並傳入 `auth.Accounts` 能以使用多組帳號與密碼，只要符合其中一組即為通過。

```go
func main() {
	m := mego.New()
	// 所有路由都必須符合其中一組帳號與密碼。
	m.Use(auth.NewAccounts(auth.Accounts{
		"myUsername": "myPassword",
		"myAdmin":    "mySecret",
	}))
	m.Run()
}
```

## 自訂驗證

透過 `auth.NewFunc`，開發者能使用自己的驗證函式來檢驗接收到的帳號與密碼是否正確。

```go
func main() {
	m := mego.New()
	// 開發者可自訂驗證函式來確認客戶端發送的帳號與密碼是否正確，
	// 這很適合與資料庫相互連接。
	m.Use(auth.NewFunc(func(username, password string) bool {
		return username == "myUsername" && password == "myPassword"
	}))
	m.Run()
}
```