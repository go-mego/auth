package main

import (
	"net/http"

	"github.com/go-mego/auth"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.GET("/", auth.NewAccounts(&auth.Accounts{
		"admin":      "admin",
		"yamiodymel": "yamiodymel",
	}), func(c *mego.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})
	e.Run()
}
