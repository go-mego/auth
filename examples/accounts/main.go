package main

import (
	"github.com/go-mego/auth"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.GET("/", auth.NewAccounts(&auth.Accounts{
		"admin":      "admin",
		"yamiodymel": "yamiodymel",
	}), func(c *mego.Context) {
		c.String(200, "Hello, world!")
	})
	e.Run()
}
