package main

import (
	"github.com/go-mego/auth"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.GET("/", auth.NewFunc(func(user, pass string) bool {
		// 這有可能造成時間差攻擊，正式場合請使用 `crypto/subtle`。
		return user == "admin" && pass == "admin"
	}), func(c *mego.Context) {
		c.String(200, "Hello, world!")
	})
	e.Run()
}
